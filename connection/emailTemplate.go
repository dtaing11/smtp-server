package connection

import "fmt"

func emailTemplate(lastName string) string {
	if lastName == "" {
		lastName = "there"
	}

	return fmt.Sprintf(`
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Thanks for reaching out</title>
  </head>
  <body style="margin:0; padding:0; background:#0f172a; font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Arial,sans-serif; color:#e5e7eb;">
    <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" style="background:#0f172a; padding:28px 12px;">
      <tr>
        <td align="center">
          <table role="presentation" width="600" cellpadding="0" cellspacing="0" style="width:100%%; max-width:600px;">

            <!-- Header -->
            <tr>
              <td style="padding:6px 4px 16px 4px;">
                <div style="font-size:13px; color:#9ca3af;">
                  Auto-reply • Dina Taing
                </div>
              </td>
            </tr>

            <!-- Card -->
            <tr>
              <td style="background:#020617; border-radius:18px; box-shadow:0 20px 50px rgba(0,0,0,0.45); overflow:hidden; border:1px solid #1e293b;">
                
                <!-- Gradient accent -->
                <div style="height:6px; background:linear-gradient(90deg,#8b5cf6,#f59e0b);"></div>

                <table width="100%%" cellpadding="0" cellspacing="0">
                  <tr>
                    <td style="padding:30px;">
                      <h1 style="margin:0; font-size:22px; color:#f8fafc;">
                        Thanks for reaching out, %s
                      </h1>

                      <p style="margin:14px 0; font-size:15px; line-height:1.65; color:#cbd5f5;">
                        I’ve received your message and appreciate you taking the time to connect.
                        I’ll follow up as soon as possible.
                      </p>

                      <!-- Info box -->
                      <div style="margin-top:18px; background:#020617; border:1px solid #334155; border-radius:14px; padding:16px;">
                        <p style="margin:0; font-size:14px; line-height:1.6; color:#e5e7eb;">
                          <strong style="color:#fbbf24;">For hiring or recruiting</strong><br/>
                          Including the role, company, timeline, and interview availability
                          (with timezone) helps me respond more efficiently.
                        </p>
                      </div>

                      <p style="margin-top:22px; font-size:14px; color:#e5e7eb;">
                        Best regards,<br/>
                        <strong style="color:#f8fafc;">Dina Taing</strong><br/>
                        <span style="color:#94a3b8;">Software Engineer</span>
                      </p>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>

            <!-- Footer -->
            <tr>
              <td style="padding:18px; text-align:center; font-size:12px; color:#64748b;">
                © 2025 Dina Taing<br/>
                Delivered via a self-made SMTP server (Go) hosted on GCP Cloud Run
              </td>
            </tr>

          </table>
        </td>
      </tr>
    </table>
  </body>
</html>
`, lastName)
}
