# Security Policy

## Supported Versions

We provide security updates for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of our software seriously. If you believe you have found a security vulnerability in the OpenTelemetry Plugin for Genkit Go, please report it to us as described below.

### How to Report a Security Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please send an email to security@xavidop.me with the following information:

1. **Type of issue** (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
2. **Full paths** of source file(s) related to the manifestation of the issue
3. **The location** of the affected source code (tag/branch/commit or direct URL)
4. **Any special configuration** required to reproduce the issue
5. **Step-by-step instructions** to reproduce the issue
6. **Proof-of-concept or exploit code** (if possible)
7. **Impact** of the issue, including how an attacker might exploit the issue

This information will help us triage your report more quickly.

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours.
- **Initial Assessment**: We will provide an initial assessment of the report within 5 business days.
- **Regular Updates**: We will keep you informed of our progress towards resolving the issue.
- **Resolution**: We aim to resolve critical security issues within 30 days of the initial report.

### Disclosure Policy

- We ask that you give us a reasonable amount of time to address the issue before any disclosure to the public or a third party.
- We will credit you for the discovery (unless you prefer to remain anonymous).
- We may ask you to test our fix to ensure the vulnerability has been properly addressed.

### Security Best Practices

When using this plugin, please follow these security best practices:


#### Input Validation

- Always validate and sanitize user inputs before passing them to AI models
- Be aware of potential prompt injection attacks
- Implement rate limiting for production applications
- Monitor and log AI model interactions for security analysis

#### Data Privacy

- Be mindful of sensitive data being sent to AI models
- Consider data residency requirements
- Implement proper data retention policies
- Follow your organization's data governance policies

### Vulnerability Response Team

Our security response team consists of:

- Xavier Portilla Edo (@xavidop) - Project Maintainer

### Scope

This security policy applies to:

- The core plugin code (`opentelemetry.go`)
- All example applications
- Documentation that could impact security
- Dependencies and their configurations

### Out of Scope

The following are considered out of scope for this security policy:

- Issues in dependencies that are already publicly known
- Issues that require physical access to a user's machine
- Social engineering attacks

### Legal

We are committed to working with security researchers and will not pursue legal action against researchers who:

- Report vulnerabilities in accordance with this policy
- Avoid violating the privacy of others
- Avoid destroying data or degrading services
- Do not access or modify data that doesn't belong to them

Thank you for helping to keep our software and community safe!
