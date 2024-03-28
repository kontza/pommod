package cmd

import (
	"github.com/beevik/etree"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type Parent struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func rootRunner(cmd *cobra.Command, args []string) {
	pom := etree.NewDocument()
	if err := pom.ReadFromFile(args[0]); err != nil {
		log.Panic().Err(err).Msg("File read failed")
		return
	}
	for _, project := range pom.SelectElements("project") {
		for _, parent := range project.SelectElements("parent") {
			for _, child := range parent.ChildElements() {
				switch child.Tag {
				case "groupId":
					child.SetText(groupId)
					break
				case "artifactId":
					child.SetText(artifactId)
					break
				case "version":
					child.SetText(version)
					break
				case "relativePath":
					parent.RemoveChild(child)
					break
				}
			}
			parent.ReindexChildren()
		}
	}

	err := pom.WriteToFile(args[0])
	if err != nil {
		log.Error().Err(err).Msg("WriteFile failed")
	}

	log.Info().Msg("All done, TTFN.")
}
