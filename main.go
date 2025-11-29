package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AljaBurgar/Redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

var studenti = map[string]redovalnica.Student{
	"63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
	"63210002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8, 6, 7, 5, 8}},
	"63210003": {Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3}},
}

func main() {

	app := &cli.Command{
		Name:  "redovalnica",
		Usage: "Upravljanje ocen študentov",

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcena",
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "najmanjša mogoča ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "največja mogoča ocena",
				Value: 10,
			},
		},

		Action: func(ctx context.Context, app *cli.Command) error {
			fmt.Println("=== REDOVALNICA ===")

			min := app.Int("minOcena")
			max := app.Int("maxOcena")

			fmt.Printf("%d %d\n", min, max)

			fmt.Println("Dodajanje ocene:")
			redovalnica.DodajOceno(studenti, "63210001", 10)
			redovalnica.DodajOceno(studenti, "63210001", 11)

			fmt.Println()
			redovalnica.IzpisRedovalnice(studenti)

			fmt.Println()
			redovalnica.IzpisiKoncniUspeh(studenti)

			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
