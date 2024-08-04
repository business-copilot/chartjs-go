package chart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDataset(t *testing.T) {
	label := "Test Dataset"
	data := []float64{10, 20, 30}
	bgColor := []string{"rgba(255, 99, 132, 0.2)"}
	borderColor := []string{"rgba(255, 99, 132, 1)"}
	borderWidth := 2

	dataset := NewDataset(label, data, bgColor, borderColor, borderWidth)

	require.Equal(t, label, dataset.Label)
	require.Equal(t, data, dataset.Data)
	require.Equal(t, bgColor, dataset.BackgroundColor)
	require.Equal(t, borderColor, dataset.BorderColor)
	require.Equal(t, borderWidth, dataset.BorderWidth)
}

func TestAddData(t *testing.T) {
	data := []float64{10, 20, 30}
	dataset := NewDataset("Test Dataset", data, nil, nil, 1)

	err := dataset.AddData([]float64{40, 50})
	require.NoError(t, err)
	require.Equal(t, []float64{10, 20, 30, 40, 50}, dataset.Data)

	err = dataset.AddData([]int{60, 70})
	require.NoError(t, err)
	require.Equal(t, []float64{10, 20, 30, 40, 50, 60, 70}, dataset.Data)

	err = dataset.AddData("unsupported type")
	require.Error(t, err)
}

func TestNewChart(t *testing.T) {
	labels := []string{"January", "February", "March"}
	data := []float64{100, 200, 300}
	dataset := NewDataset("Test Dataset", data, nil, nil, 1)
	chartData := ChartData{
		Labels:   labels,
		Datasets: []Dataset{dataset},
	}
	responsive := true
	options := Options{
		Responsive: &responsive,
	}

	chart := NewChart("bar", chartData, options)

	require.Equal(t, "bar", chart.Type)
	require.Equal(t, labels, chart.Data.Labels)
	require.Equal(t, data, chart.Data.Datasets[0].Data)
	require.True(t, *chart.Options.Responsive)
}

func TestUpdateData(t *testing.T) {
	chart := NewChart("line", ChartData{}, Options{})

	newLabels := []string{"April", "May"}
	newData := []float64{150, 250}
	newDataset := NewDataset("Updated Dataset", newData, nil, nil, 1)
	newChartData := ChartData{
		Labels:   newLabels,
		Datasets: []Dataset{newDataset},
	}

	chart.UpdateData(newChartData)

	require.Equal(t, newLabels, chart.Data.Labels)
	require.Equal(t, newData, chart.Data.Datasets[0].Data)
}

func TestAddDataToDataset(t *testing.T) {
	labels := []string{"January", "February"}
	data := []float64{100, 200}
	dataset := NewDataset("Dataset 1", data, nil, nil, 1)
	chartData := ChartData{
		Labels:   labels,
		Datasets: []Dataset{dataset},
	}
	responsive := true
	options := Options{
		Responsive: &responsive,
	}

	chart := NewChart("bar", chartData, options)

	err := chart.AddDataToDataset("Dataset 1", []float64{300, 400})
	require.NoError(t, err)
	require.Equal(t, []float64{100, 200, 300, 400}, chart.Data.Datasets[0].Data)

	err = chart.AddDataToDataset("Nonexistent Dataset", []float64{500})
	require.Error(t, err)
}

func TestToJSON(t *testing.T) {
	labels := []string{"January", "February"}
	data := []float64{100, 200}
	dataset := NewDataset("Dataset 1", data, nil, nil, 1)
	chartData := ChartData{
		Labels:   labels,
		Datasets: []Dataset{dataset},
	}
	responsive := true
	options := Options{
		Responsive: &responsive,
	}

	chart := NewChart("bar", chartData, options)

	jsonStr, err := chart.ToJSON()
	require.NoError(t, err)
	require.JSONEq(t, `{
		"type": "bar",
		"data": {
			"labels": ["January", "February"],
			"datasets": [
				{
					"label": "Dataset 1",
					"data": [100, 200],
					"backgroundColor": null,
					"borderColor": null,
					"borderWidth": 1
				}
			]
		},
		"options": {
			"responsive": true
		}
	}`, jsonStr)
}

func TestFromJSON(t *testing.T) {
	jsonStr := `{
		"type": "bar",
		"data": {
			"labels": ["January", "February"],
			"datasets": [
				{
					"label": "Dataset 1",
					"data": [100, 200],
					"backgroundColor": null,
					"borderColor": null,
					"borderWidth": 1
				}
			]
		},
		"options": {
			"responsive": true
		}
	}`

	chart, err := FromJSON(jsonStr)
	require.NoError(t, err)

	require.Equal(t, "bar", chart.Type)
	require.Equal(t, []string{"January", "February"}, chart.Data.Labels)
	require.Equal(t, []float64{100, 200}, chart.Data.Datasets[0].Data)
	require.True(t, *chart.Options.Responsive)
}
