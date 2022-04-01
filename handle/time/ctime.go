package time

import(
	"fmt"
  "time"
	"bytes"
)

func FormatDateTime(t time.Time) string {
	var buffer bytes.Buffer
	buffer.WriteString(t.Month().String()[:3])
	buffer.WriteString(fmt.Sprintf(" %2d '%2d at %2d:%2d", t.Day(), t.Year()%100, t.Hour(), t.Minute()))
	return buffer.String()
}
func FormatIso88601(t string) string {
  t , err := time.Parse("2006/01/02 15:04:05",t)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
  return t.Format("2006-01-02T15:04:05")
} 