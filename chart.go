package chart

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DataPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type Dataset struct {
	Label           string    `json:"label"`
	Data            []float64 `json:"data"`
	BackgroundColor []string  `json:"backgroundColor"`
	BorderColor     []string  `json:"borderColor"`
	BorderWidth     int       `json:"borderWidth"`
}

func NewDataset(label string, data []float64, bgColor []string, borderColor []string, borderWidth int) Dataset {
	return Dataset{
		Label:           label,
		Data:            data,
		BackgroundColor: bgColor,
		BorderColor:     borderColor,
		BorderWidth:     borderWidth,
	}
}

func (d *Dataset) AddData(data interface{}) error {
	switch v := data.(type) {
	case []float64:
		d.Data = append(d.Data, v...)
	case []float32:
		for _, val := range v {
			d.Data = append(d.Data, float64(val))
		}
	case []int:
		for _, val := range v {
			d.Data = append(d.Data, float64(val))
		}
	case float64:
		d.Data = append(d.Data, v)
	case float32:
		d.Data = append(d.Data, float64(v))
	case int:
		d.Data = append(d.Data, float64(v))
	default:
		return errors.New("unsupported data type")
	}
	return nil
}

type ChartData struct {
	Labels   []string  `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Chart struct {
	Type    string    `json:"type"`
	Data    ChartData `json:"data"`
	Options Options   `json:"options"`
}

func NewChart(chartType string, data ChartData, options Options) Chart {
	return Chart{
		Type:    chartType,
		Data:    data,
		Options: options,
	}
}

func (c *Chart) UpdateData(newData ChartData) {
	c.Data = newData
}

func (c *Chart) AddDataToDataset(label string, data interface{}) error {
	for i, dataset := range c.Data.Datasets {
		if dataset.Label == label {
			err := c.Data.Datasets[i].AddData(data)
			if err != nil {
				return fmt.Errorf("error adding data to dataset: %v", err)
			}
			return nil
		}
	}
	return fmt.Errorf("dataset with label %s not found", label)
}

func (c *Chart) ToJSON() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("error marshaling chart to JSON: %v", err)
	}
	return string(bytes), nil
}

func FromJSON(jsonStr string) (Chart, error) {
	var chart Chart
	err := json.Unmarshal([]byte(jsonStr), &chart)
	if err != nil {
		return chart, fmt.Errorf("error unmarshaling JSON to chart: %v", err)
	}
	return chart, nil
}
