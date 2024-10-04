package services

import (
	"os"
	"fmt"
	// "bytes"
	// "html/template"
	// "strings"

	// "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PageData struct {
    Name string
}

func GeneratePDF(name string, email string) error {
    // pdfg, err := wkhtmltopdf.NewPDFGenerator()
    // if err != nil {
    //     return fmt.Errorf("could not create PDF generator: %v", err)
    // }
    // tmpl, err := template.ParseFiles("assets/pdf.html")
    // if err != nil {
    //     return fmt.Errorf("could not parse template: %v", err)
    // }

    // var htmlContent string
    // buf := new(bytes.Buffer)
    // err = tmpl.Execute(buf, PageData{Name: name})
    // if err != nil {
    //     return fmt.Errorf("could not execute template: %v", err)
    // }
    // htmlContent = buf.String()

    // pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent)))

    // err = pdfg.Create()
    // if err != nil {
    //     return fmt.Errorf("could not create PDF: %v", err)
    // }

    // err = pdfg.WriteFile("output.pdf")
    // if err != nil {
    //     return fmt.Errorf("could not save PDF file: %v", err)
    // }
    fmt.Println("hello there 11")

    // pdfFilePath := "test/report.pdf"
    smtpPort := 587
    // from := os.Getenv("SMTP_FROM")
    smtpServer := os.Getenv("SMTP_SERVER")
    smtpUser := os.Getenv("SMTP_USERNAME")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    // err := SendEmailWithAttachment(subject, to, from, smtpServer, smtpPort, username, password)
    // if err != nil {
    //     fmt.Println("Error sending email:", err)
    // } else {
    //     fmt.Println("Email sent successfully!")
    // }


	// Email content
	to := email
	subject := "Final Inspection Report"
	body := `<p>Hello Inspector,</p>
<p>Here is your final CAT vechical inspection report </p>
<p style="font-size:0;line-height:0;">
  <a href="https://drive.google.com/file/d/1YztrIGLJlx-LzV1o0vzYED2uMxKIbSCI/view?usp=sharing">
    Hidden Google Drive Link
  </a>
</p>
<p>Thank you!</p>`

	// Send the email
	err := SendEmail(to, subject, body, smtpServer, smtpPort, smtpUser, smtpPassword)
	if err != nil {
		fmt.Println("Error sending email: %v", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

    return nil
}

