package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func downFile(main_url, url_path string) error {

	http_response, err_req := http.Get(main_url + url_path)

	if err_req != nil {
		return err_req
	}

	file_content, err_read := ioutil.ReadAll(http_response.Body)

	if err_read != nil {
		return err_read
	}

	return ioutil.WriteFile(url_path, file_content, os.ModePerm)

}

func main() {
	var url = "http://www.kdnuggets.com/data_mining_course/"
	var urlPaths = []string{
		"dm1-introduction-ml-data-mining.ppt",
		//DM2: Machine Learning and Classification, updated June 7, 2006.
		"dm2-intro-machine-learning-classification.ppt",
		//DM3: Input: Concepts, Instances, Attributes.
		"dm3-input-concepts.ppt",
		//DM4: Output: Knowledge Representation, updated June 7, 2006.
		"dm4-output-representation.ppt",
		//DM5: Classification - Basic Methods.
		"dm5-classification-basic.ppt",
		//DM6: DM6: Classification: Decision Trees.
		"dm6-decision-tree-intro.ppt",
		//DM7: Classification: C4.5.
		"dm7-decision-tree-c45.ppt",
		//DM8: Classification: CART.
		"dm8-decision-tree-cart.ppt",
		//DM9: Classification: Rules, Regression, K-Nearest Neighbour.
		"dm9-rules-regression-knn.ppt",
		//DM10: Evaluation and Credibility, updated May 31, 2006.
		"dm10-evaluation.ppt",
		//DM11: Evaluation - Lift and Costs, updated May 31, 2006.
		"dm11-evaluation-lift-cost.ppt",
		//DM12: Data Preparation for Knowledge Discovery, updated June 7, 2006.
		"dm12-data-preparation.ppt",
		//DM13: Clustering, updated May 31, 2006.
		"dm13-clustering.ppt",
		//DM14: Associations Rules, updated May 31, 2006.
		"dm14-association-rules.ppt",
		//DM15: Data Mining and Visualization, (3.2MB), updated May 31, 2006.
		"dm15-visualization-data-mining.ppt",
		//DM16: Summarization and Deviation Detection.
		"dm16-summarization-deviation-detection.ppt",
		//DM17: Applications: Targeted Marketing, KDD Cup, and Customer Modeling, updated Oct 18, 2004.
		"dm17-targeted-marketing-kdd-cup.ppt",
		//DM18: Applications: Genomic Microarray Data Mining.
		"dm18-microarray-data-mining.ppt",
		//DM19: Data Mining and Society; Future Directions
		"dm19-data-mining-and-society.ppt",
	}

	for _, curr_url := range urlPaths {
		fmt.Println("Starting downloading: ", curr_url)
		if err_down := downFile(url, curr_url); err_down != nil {
			fmt.Println("Error downloading file: ", err_down)
		}

	}

}
