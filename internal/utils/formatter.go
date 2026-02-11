package utils

import (
	"fmt"
	"strings"
	"time"

	"tls-checker/internal/model"
)

// Format timestamp in milliseconds to a human-readable date
func formatTimestamp(ms int64) string {
	if ms == 0 {
		return "N/A"
	}
	t := time.UnixMilli(ms)
	return t.Format("Jan 02, 2006 15:04 MST")
}

// Format duration in milliseconds
func formatDuration(ms int) string {
	duration := time.Duration(ms) * time.Millisecond
	if duration < time.Second {
		return fmt.Sprintf("%dms", ms)
	}
	return fmt.Sprintf("%.1fs", duration.Seconds())
}

// Section header
func sectionHeader(title string) string {
	return fmt.Sprintf("\n[::b][%s]━━━ %s ━━━[-][::B]\n", TagPrimary, title)
}

// Subsection header
func subHeader(title string) string {
	return fmt.Sprintf("\n  [::b][%s]── %s ──[-][::B]\n", TagText, title)
}

// Key-value row with alignment
func row(key, value string) string {
	return fmt.Sprintf("  [%s]%-28s[-] %s\n", TagLabel, key, value)
}

// Boolean status with colored indicator
func boolStatus(val bool, trueText, falseText string) string {
	if val {
		return fmt.Sprintf("[%s]✔ %s[-]", TagOk, trueText)
	}
	return fmt.Sprintf("[%s]✘ %s[-]", TagError, falseText)
}

// Colored yes/no
func yesNo(val bool) string {
	return boolStatus(val, "Yes", "No")
}

// FormatEndpoint produces a formatted string for an Endpoint
func FormatEndpoint(endpoint *model.Endpoint) string {
	var builder strings.Builder

	// Summary
	writeSummary(&builder, endpoint)

	if endpoint.Details.Cert.Subject != "" {
		// Certificate
		writeCertificate(&builder, endpoint.Details.Cert)

		// Certificate Chain
		writeChain(&builder, endpoint.Details.Chain)
	}

	// Protocols
	writeProtocols(&builder, endpoint.Details.Protocols)

	// Cipher Suites
	writeSuites(&builder, endpoint.Details.Suites)

	// Handshake Simulation
	writeSimulations(&builder, endpoint.Details.Sims)

	// Server Configuration
	writeServerConfig(&builder, endpoint.Details)

	// Vulnerabilities
	writeVulnerabilities(&builder, endpoint.Details)

	// HTTP Security Headers
	writeHTTPHeaders(&builder, endpoint.Details)

	return builder.String()
}

// Summary
func writeSummary(builder *strings.Builder, endpoint *model.Endpoint) {

	color := gradeColor(endpoint.Grade)
	fmt.Fprintf(builder, "  [%s::b]  ╔═══════════╗[-::-]\n", color)
	fmt.Fprintf(builder, "  [%s::b]  ║           ║[-::-]\n", color)
	fmt.Fprintf(builder, "  [%s::b]  ║    %s    ║[-::-]\n", color, centerGrade(endpoint.Grade))
	fmt.Fprintf(builder, "  [%s::b]  ║           ║[-::-]\n", color)
	fmt.Fprintf(builder, "  [%s::b]  ╚═══════════╝[-::-]\n\n", color)

	builder.WriteString(row("IP Address", endpoint.IPAddress))
	builder.WriteString(row("Status", endpoint.StatusMessage))
	builder.WriteString(row("Grade", fmt.Sprintf("[%s::b]%s[-::-]", color, endpoint.Grade)))
	if endpoint.GradeTrustIgnored != endpoint.Grade {
		builder.WriteString(row("Grade (trust ignored)", fmt.Sprintf("[%s]%s[-]", gradeColor(endpoint.GradeTrustIgnored), endpoint.GradeTrustIgnored)))
	}
	builder.WriteString(row("Has Warnings", yesNo(endpoint.HasWarnings)))
	builder.WriteString(row("Exceptional", yesNo(endpoint.IsExceptional)))
	builder.WriteString(row("Scan Duration", formatDuration(endpoint.Duration)))
	builder.WriteString(row("Progress", fmt.Sprintf("%d%%", endpoint.Progress)))

	if endpoint.Details.ServerSignature != "" {
		builder.WriteString(row("Server Signature", endpoint.Details.ServerSignature))
	}

	// Key info summary
	if endpoint.Details.Key.Alg != "" {
		builder.WriteString(row("Server Key", fmt.Sprintf("%s %d bits (Strength: %d)", endpoint.Details.Key.Alg, endpoint.Details.Key.Size, endpoint.Details.Key.Strength)))
	}
}

// Center security grade on screen
func centerGrade(grade string) string {
	switch len(grade) {
	case 1:
		return " " + grade + " "
	case 2:
		return " " + grade + ""
	case 3:
		return grade + " "
	default:
		return grade
	}
}

// Certificate
func writeCertificate(builder *strings.Builder, cert model.Cert) {
	title := fmt.Sprintf("Certificate: %s (%s)", cert.CommonNames[0], cert.SigAlg)
	builder.WriteString(sectionHeader(title))

	builder.WriteString(row("Subject", cert.Subject))
	builder.WriteString(row("Common Names", strings.Join(cert.CommonNames, ", ")))

	if len(cert.AltNames) > 0 {
		// Show first few alt names, then count remaining
		displayNames := cert.AltNames
		suffix := ""
		if len(displayNames) > 5 {
			suffix = fmt.Sprintf(" [%s](+%d more)[-]", TagLabel, len(displayNames)-5)
			displayNames = displayNames[:5]
		}
		builder.WriteString(row("Alternative Names", strings.Join(displayNames, ", ")+suffix))
	}

	builder.WriteString(row("Valid From", formatTimestamp(cert.NotBefore)))
	builder.WriteString(row("Valid Until", formatTimestamp(cert.NotAfter)))

	// Check if expired or near expiry
	notAfter := time.UnixMilli(cert.NotAfter)
	now := time.Now()
	if now.After(notAfter) {
		builder.WriteString(row("", fmt.Sprintf("[%s::b]⚠ CERTIFICATE EXPIRED[-::-]", TagError)))
	} else if notAfter.Sub(now) < 30*24*time.Hour {
		daysLeft := int(notAfter.Sub(now).Hours() / 24)
		builder.WriteString(row("", fmt.Sprintf("[%s::b]⚠ Expires in %d days[-::-]", TagWarning, daysLeft)))
	}

	builder.WriteString(row("Issuer", cert.IssuerLabel))
	builder.WriteString(row("Signature Algorithm", cert.SigAlg))
	builder.WriteString(row("SHA1 Fingerprint", cert.Sha1Hash))
	builder.WriteString(row("Pin SHA256", cert.PinSha256))

	if cert.Sct {
		builder.WriteString(row("SCT", fmt.Sprintf("[%s]Yes[-]", TagOk)))
	} else {
		builder.WriteString(row("SCT", fmt.Sprintf("[%s]No[-]", TagError)))
	}

	revStatus := revocationStatusText(cert.RevocationStatus)
	builder.WriteString(row("Revocation Status", revStatus))

	if len(cert.OcspURIs) > 0 {
		builder.WriteString(row("OCSP URI", strings.Join(cert.OcspURIs, ", ")))
	}
	if len(cert.CrlURIs) > 0 {
		builder.WriteString(row("CRL URI", strings.Join(cert.CrlURIs, ", ")))
	}
}

// Format for revocation status attribute
func revocationStatusText(status int) string {
	switch status {
	case 0:
		return fmt.Sprintf("[%s]Not checked[-]", TagWarning)
	case 1:
		return fmt.Sprintf("[%s]Revoked[-]", TagError)
	case 2:
		return fmt.Sprintf("[%s]Not revoked[-]", TagOk)
	case 3:
		return fmt.Sprintf("[%s]Revocation check error[-]", TagWarning)
	case 4:
		return fmt.Sprintf("[%s]No revocation information[-]", TagWarning)
	default:
		return fmt.Sprintf("Unknown (%d)", status)
	}
}

// Certificate Chain
func writeChain(builder *strings.Builder, chain model.Chain) {
	if len(chain.Certs) == 0 {
		return
	}

	builder.WriteString(sectionHeader("Certificate Chain"))

	for i, cert := range chain.Certs {
		builder.WriteString(subHeader(fmt.Sprintf("#%d: %s", i+1, cert.Label)))
		builder.WriteString(row("Subject", cert.Subject))
		builder.WriteString(row("Issuer", cert.IssuerLabel))
		builder.WriteString(row("Valid From", formatTimestamp(cert.NotBefore)))
		builder.WriteString(row("Valid Until", formatTimestamp(cert.NotAfter)))
		builder.WriteString(row("Key", fmt.Sprintf("%s %d bits", cert.KeyAlg, cert.KeySize)))
		builder.WriteString(row("Signature Algorithm", cert.SigAlg))
		builder.WriteString(row("Fingerprint (SHA1)", cert.Sha1Hash))
	}

	if chain.Issues != 0 {
		fmt.Fprintf(builder, "\n  [%s]Chain Issues: %d[-]\n", TagError, chain.Issues)
	}
}

// Protocols
func writeProtocols(builder *strings.Builder, protocols []model.Protocols) {
	if len(protocols) == 0 {
		return
	}

	builder.WriteString(sectionHeader("Protocols"))

	for _, p := range protocols {
		icon := fmt.Sprintf("[%s]✔[-]", TagOk)
		// Flag old protocols
		if p.Name == "SSL" || (p.Name == "TLS" && (p.Version == "1.0" || p.Version == "1.1")) {
			icon = fmt.Sprintf("[%s]⚠[-]", TagWarning)
		}
		fmt.Fprintf(builder, "  %s  %-6s %s\n", icon, p.Name, p.Version)
	}
}

// Cipher Suites
func writeSuites(builder *strings.Builder, suites model.Suites) {
	if len(suites.List) == 0 {
		return
	}

	builder.WriteString(sectionHeader("Cipher Suites"))

	if suites.Preference {
		fmt.Fprintf(builder, "  [%s]Server has cipher order preference[-]\n\n", TagLabel)
	}

	fmt.Fprintf(builder, "  [::b]%-50s %s    %s[::B]\n", "Cipher", "Bits", "Key Exchange")
	fmt.Fprintf(builder, "  %s\n", strings.Repeat("─", 80))

	for _, s := range suites.List {
		strength := suiteStrengthColor(s.CipherStrength)
		kx := suiteKxInfo(s)
		fmt.Fprintf(builder, "  %-50s %s%-4d[-]    %s\n", s.Name, strength, s.CipherStrength, kx)
	}
}

func suiteKxInfo(suite model.Suite) string {
	if suite.EcdhBits > 0 {
		return fmt.Sprintf("ECDH %d bits (strength %d)", suite.EcdhBits, suite.EcdhStrength)
	}
	if suite.DhStrength > 0 {
		return fmt.Sprintf("DH %d bits", suite.DhStrength)
	}
	return ""
}

// Handshake Simulation
func writeSimulations(builder *strings.Builder, sims model.Sims) {
	if len(sims.Results) == 0 {
		return
	}

	builder.WriteString(sectionHeader("Handshake Simulation"))
	fmt.Fprintf(builder, "  [::b]%-40s %-10s %s[::B]\n", "Client", "Protocol", "Cipher Suite")
	fmt.Fprintf(builder, "  %s\n", strings.Repeat("─", 90))

	for _, sim := range sims.Results {
		clientName := formatSimClient(sim.Client)
		if sim.ErrorCode != 0 {
			fmt.Fprintf(builder, "  %-40s [%s]Failed (error %d)[-]\n", clientName, TagError, sim.ErrorCode)
		} else {
			protocol := protocolName(sim.ProtocolId)
			suite := fmt.Sprintf("0x%04X", sim.SuiteId)
			kx := ""
			if sim.KxInfo != "" {
				kx = fmt.Sprintf(" [%s](%s)[-]", TagLabel, sim.KxInfo)
			}
			fmt.Fprintf(builder, "  %-40s [%s]%-10s[-] %s%s\n", clientName, TagOk, protocol, suite, kx)
		}
	}
}

// Returns simulation client formatted string
func formatSimClient(client map[string]any) string {
	name, _ := client["name"].(string)
	version, _ := client["version"].(string)
	platform, _ := client["platform"].(string)

	result := name
	if version != "" {
		result += " " + version
	}
	if platform != "" {
		result += " / " + platform
	}
	return result
}

// Returns the protocol name based on its id
func protocolName(id int) string {
	switch id {
	case 769:
		return "TLS 1.0"
	case 770:
		return "TLS 1.1"
	case 771:
		return "TLS 1.2"
	case 772:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("0x%04X", id)
	}
}

// Server Configuration
func writeServerConfig(builder *strings.Builder, d model.Details) {
	builder.WriteString(sectionHeader("Server Configuration"))

	builder.WriteString(row("Forward Secrecy", forwardSecrecyText(d.ForwardSecrecy)))
	builder.WriteString(row("ALPN Support", yesNo(d.SupportsAlpn)))
	builder.WriteString(row("NPN Support", yesNo(d.SupportsNpn)))
	builder.WriteString(row("Session Resumption", sessionResumptionText(d.SessionResumption)))
	builder.WriteString(row("Session Tickets", sessionTicketText(d.SessionTickets)))
	builder.WriteString(row("OCSP Stapling", yesNo(d.OcspStapling)))
	builder.WriteString(row("SNI Required", yesNo(d.SniRequired)))
	builder.WriteString(row("HTTP Status Code", fmt.Sprintf("%d", d.HTTPStatusCode)))
	builder.WriteString(row("Renegotiation Support", renegotiationText(d.RenegSupport)))
	builder.WriteString(row("Compression", compressionText(d.CompressionMethods)))
}

func forwardSecrecyText(forwardSecrecy int) string {
	switch forwardSecrecy {
	case 0:
		return fmt.Sprintf("[%s]Not supported[-]", TagError)
	case 1:
		return fmt.Sprintf("[%s]With some browsers[-]", TagWarning)
	case 2:
		return fmt.Sprintf("[%s]With modern browsers[-]", TagWarning)
	case 4:
		return fmt.Sprintf("[%s]Yes (with most browsers) ROBUST[-]", TagOk)
	default:
		return fmt.Sprintf("%d", forwardSecrecy)
	}
}

func sessionResumptionText(sessionResumption int) string {
	switch sessionResumption {
	case 0:
		return "Not supported"
	case 1:
		return "No (IDs assigned but not accepted)"
	case 2:
		return fmt.Sprintf("[%s]Yes[-]", TagOk)
	default:
		return fmt.Sprintf("%d", sessionResumption)
	}
}

func sessionTicketText(sessionTicket int) string {
	switch sessionTicket {
	case 0:
		return "Not supported"
	case 1:
		return fmt.Sprintf("[%s]Supported[-]", TagOk)
	case 2:
		return fmt.Sprintf("[%s]Implementation is faulty[-]", TagOk)
	default:
		return fmt.Sprintf("%d", sessionTicket)
	}
}

func renegotiationText(renegotiationText int) string {
	switch renegotiationText {
	case 0:
		return fmt.Sprintf("[%s]Not supported[-]", TagError)
	case 1:
		return fmt.Sprintf("[%s]Insecure client-initiated[-]", TagWarning)
	case 2:
		return fmt.Sprintf("[%s]Secure renegotiation supported[-]", TagOk)
	default:
		return fmt.Sprintf("%d", renegotiationText)
	}
}

func compressionText(compressionText int) string {
	if compressionText == 0 {
		return fmt.Sprintf("[%s]No (good)[-]", TagOk)
	}
	return fmt.Sprintf("[%s]Yes (%d methods) — CRIME vulnerable[-]", TagError, compressionText)
}

// Vulnerabilities
func writeVulnerabilities(builder *strings.Builder, details model.Details) {
	builder.WriteString(sectionHeader("Known Vulnerabilities"))

	writeVulnerabilityRow(builder, "Heartbleed (CVE-2014-0160)", !details.Heartbleed)
	writeVulnerabilityRow(builder, "POODLE (SSLv3)", !details.Poodle)
	writeVulnerabilityRow(builder, "POODLE (TLS)", details.PoodleTls != 2)
	writeVulnerabilityRow(builder, "Downgrade (TLS_FALLBACK_SCSV)", details.FallbackScsv)
	writeVulnerabilityRow(builder, "FREAK", !details.Freak)
	writeVulnerabilityRow(builder, "Logjam", !details.Logjam)
	writeVulnerabilityRow(builder, "BEAST", !details.VulnBeast)
	writeVulnerabilityRow(builder, "RC4 Support", !details.SupportsRc4)

	builder.WriteString(row("OpenSSL CCS (CVE-2014-0224)", openSslCcsText(details.OpenSslCcs)))
	builder.WriteString(row("OpenSSL Lucky13", openSslLucky13Text(details.OpenSSLLuckyMinus20)))
}

func writeVulnerabilityRow(builder *strings.Builder, name string, safe bool) {
	if safe {
		builder.WriteString(row(name, fmt.Sprintf("[%s]No[-]", TagOk)))
	} else {
		builder.WriteString(row(name, fmt.Sprintf("[%s::b]VULNERABLE[-::-]", TagError)))
	}
}

func openSslCcsText(val int) string {
	switch val {
	case -1:
		return "Test failed"
	case 0:
		return "Unknown"
	case 1:
		return fmt.Sprintf("[%s]Not vulnerable[-]", TagOk)
	case 2:
		return fmt.Sprintf("[%s]Possibly vulnerable (not exploitable)[-]", TagWarning)
	case 3:
		return fmt.Sprintf("[%s::b]VULNERABLE[-::-]", TagError)
	default:
		return fmt.Sprintf("%d", val)
	}
}

func openSslLucky13Text(val int) string {
	switch val {
	case -1:
		return "Test failed"
	case 0:
		return "Unknown"
	case 1:
		return fmt.Sprintf("[%s]Not vulnerable[-]", TagOk)
	case 2:
		return fmt.Sprintf("[%s::b]VULNERABLE[-::-]", TagError)
	default:
		return fmt.Sprintf("%d", val)
	}
}

// HTTP Security Headers
func writeHTTPHeaders(builder *strings.Builder, details model.Details) {
	builder.WriteString(sectionHeader("HTTP Security Headers"))

	// HSTS
	hstsStatus := details.HstsPolicy.Status
	if hstsStatus == "present" {
		builder.WriteString(row("HSTS", fmt.Sprintf("[%s]Enabled[-] (max-age=%d)", TagOk, details.HstsPolicy.MaxAge)))
		builder.WriteString(row("HSTS Subdomains", yesNo(details.StsSubdomains)))
		builder.WriteString(row("HSTS Preload", yesNo(details.StsPreload)))

		// Check if max-age is long enough
		if details.HstsPolicy.MaxAge >= details.HstsPolicy.LongMaxAge {
			builder.WriteString(row("", fmt.Sprintf("[%s]✔ Long max-age (≥6 months)[-]", TagOk)))
		} else {
			builder.WriteString(row("", fmt.Sprintf("[%s]⚠ Short max-age (< 6 months)[-]", TagWarning)))
		}
	} else {
		builder.WriteString(row("HSTS", fmt.Sprintf("[%s]Not present[-]", TagError)))
	}

	// HSTS Preload status across browsers
	if len(details.HstsPreloads) > 0 {
		builder.WriteString(subHeader("HSTS Preload Status"))
		for _, p := range details.HstsPreloads {
			status := fmt.Sprintf("[%s]Absent[-]", TagError)
			if p.Status == "present" {
				status = fmt.Sprintf("[%s]Present[-]", TagOk)
			}
			builder.WriteString(row(p.Source, status))
		}
	}

	// HPKP
	if details.HpkpPolicy.Status == "present" {
		builder.WriteString(row("HPKP", fmt.Sprintf("[%s]Present[-]", TagOk)))
	} else {
		builder.WriteString(row("HPKP", fmt.Sprintf("[%s]Not present[-]", TagLabel)))
	}
}
