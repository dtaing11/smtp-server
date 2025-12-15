package connection

import "fmt"

func emailTemplate(lastName string) string {
	return fmt.Sprintf(`
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Thanks for reaching out</title>
  </head>
  <body style="margin:0; padding:0; background:#f6f7fb; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Arial, sans-serif; color:#111827;">
    <table role="presentation" cellpadding="0" cellspacing="0" border="0" width="100%" style="background:#f6f7fb; padding:24px 12px;">
      <tr>
        <td align="center">
          <table role="presentation" cellpadding="0" cellspacing="0" border="0" width="600" style="width:100%; max-width:600px;">

            <!-- Header -->
            <tr>
              <td style="padding: 8px 4px 16px 4px;">
                <div style="font-size:14px; color:#6b7280;">
                  Auto-reply • Dina Taing
                </div>
              </td>
            </tr>

            <!-- Card -->
            <tr>
              <td style="background:#ffffff; border-radius:16px; box-shadow: 0 8px 30px rgba(17,24,39,0.08); overflow:hidden;">
                <div style="height:6px; background: linear-gradient(90deg, #6366f1, #22c55e);"></div>

                <table width="100%%">
                  <tr>
                    <td style="padding:28px;">
                      <h1 style="margin:0; font-size:22px;">
                        Thank you for reaching out, %s.
                      </h1>
                      <p style="margin:12px 0; font-size:15px; line-height:1.6;">
                        I’ve received your message and appreciate your interest.
                        I’ll get back to you as soon as possible.
                      </p>

                      <div style="background:#f9fafb; border:1px solid #e5e7eb; border-radius:12px; padding:14px;">
                        <p style="margin:0; font-size:14px;">
                          <strong>For hiring or recruiting:</strong><br/>
                          Including the role, company, timeline, and interview availability (with timezone)
                          helps me respond more efficiently.
                        </p>
                      </div>

                      <p style="margin-top:20px; font-size:14px;">
                        Best regards,<br/>
                        <strong>Dina Taing</strong><br/>
                        <span style="color:#6b7280;">Software Engineer</span>
                      </p>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>

            <tr>
              <td style="padding:16px; text-align:center; font-size:12px; color:#9ca3af;">
                © 2025 Dina Taing
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
