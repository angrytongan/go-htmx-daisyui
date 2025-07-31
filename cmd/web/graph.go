package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const (
	// numGraphs is the number of graphs we render.
	numGraphs = 2
)

type GraphThing struct {
	Element template.HTML
	Script  template.HTML
}

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for range 7 {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}

	return items
}

func makeAllGraphs(n int) ([]GraphThing, error) {
	graphs := []GraphThing{}

	for range n {
		bar := charts.NewBar()

		bar.AddJSFuncStrs(`
			window.addEventListener('resize', () => {
				const myChart = %MY_ECHARTS%;
				myChart.resize();
			});
		`)

		bar.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{
				Height: "200px",
				Width:  "auto",
			}),
		)

		// Put data into instance
		bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
			AddSeries("Category A", generateBarItems()).
			AddSeries("Category B", generateBarItems())

		snippet := bar.RenderSnippet()

		graphs = append(graphs, GraphThing{
			Element: template.HTML(snippet.Element),
			Script:  template.HTML(snippet.Script),
		})
	}

	return graphs, nil
}

func (app *Application) widgetGraphRandom(w http.ResponseWriter, r *http.Request) {
	graphs, err := makeAllGraphs(numGraphs)

	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("makeAllGraphs(%d): %w", numGraphs, err),
			http.StatusInternalServerError,
		)

		return
	}

	pageData := map[string]any{
		"Graphs":     graphs,
		"Servertime": time.Now().String(),
	}

	app.render(w, r, "widget-graph-random", pageData, http.StatusOK)
}
