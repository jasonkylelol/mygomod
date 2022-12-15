package main

import (
	"strconv"

	pkg "github.com/apulisai/go-business/pkg/mlops"
)

func generateTrainDataset(modelDataPair *pkg.ModelDataPair, stepData *pkg.WorkflowStepData) error {
	modelDataPair.DatasetItems = []pkg.ModelDatasetItem{}
	for _, datasetItem := range stepData.Datasets {
		if modelDataPair.DatasetVersion == "etl" &&
			datasetItem.Croped == false && datasetItem.Labeled == true {
			// 没有经过抠图的数据集，是检测数据集，用于检测模型
			modelDataPair.DatasetID = strconv.Itoa(int(datasetItem.DatasetId))
			modelDataPair.DatasetVersion = datasetItem.DatasetVersion
			modelDataPair.DatasetName = datasetItem.DatasetName
			modelDataPair.DatasetItems = append(modelDataPair.DatasetItems, pkg.ModelDatasetItem{
				DatasetID:      strconv.Itoa(int(datasetItem.DatasetId)),
				DatasetVersion: datasetItem.DatasetVersion,
				DatasetName:    datasetItem.DatasetName,
			})
			// break
		} else if modelDataPair.DatasetVersion == "crop" &&
			datasetItem.Croped == true && datasetItem.Labeled == true {
			// 经过抠图的数据集，是分类数据集，用于分类模型
			modelDataPair.DatasetID = strconv.Itoa(int(datasetItem.DatasetId))
			modelDataPair.DatasetVersion = datasetItem.DatasetVersion
			modelDataPair.DatasetName = datasetItem.DatasetName
			// break
		}
	}
	return nil
}
