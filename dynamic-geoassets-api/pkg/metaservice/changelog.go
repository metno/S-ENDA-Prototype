package metaservice

import (
	"bytes"
	"html/template"
)

const dateFormat = "20019-Jan-02"

// Changelog is the structured changelog of the service
type Changelog struct {
	Title   string
	Entries []ChangelogEntry
}

// ChangelogEntry is one entry in the total Changelog to the service.
type ChangelogEntry struct {
	Title      string
	Content    string
	UpdateTime string
}

func (cl *Changelog) ToHTML() ([]byte, error) {
	t, err := template.New("Changelog").Parse(`<table>
    <tr><th width="15%">Date</th><th width="85%"> </th></tr>
	<tbody>
	  {{range .Entries}}
      <tr>
        <td>{{.UpdateTime}}</td>
        <td><h3>{{}}.Title}}</h3>
			{{}}.Content}}
		</td>
	  </tr>
	  {{end}}
    </tbody>
  </table>
`)
	if err != nil {
		return nil, err
	}

	var payload bytes.Buffer
	err = t.Execute(&payload, cl)

	if err != nil {
		return nil, err
	}

	return payload.Bytes(), nil
}

func (l *Changelog) ToAtomFeed() {

}
