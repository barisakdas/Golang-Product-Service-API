package cmd

import (
	"HB_Task/infrastructure/auth"
	"HB_Task/presentation"
	"net/http"
)

func Run() {

	http.HandleFunc("/", presentation.GetAllEndpoints)        // For endpoints
	http.HandleFunc("/get-access-token", auth.GetAccessToken) // For access token

	// Campaign endpoints
	http.Handle("/get_all_campaigns", auth.IsAuthorized(presentation.GetAllCampaigns))
	http.Handle("/get_campaign_info", auth.IsAuthorized(presentation.GetCampaignsByName))
	http.Handle("/create_campaign", auth.IsAuthorized(presentation.AddNewCampaign))
	http.Handle("/update_campaign", auth.IsAuthorized(presentation.UpdateCampaign))
	http.Handle("/increase_time", auth.IsAuthorized(presentation.IncreaseCampaignHour))
	http.Handle("/delete_campaign", auth.IsAuthorized(presentation.DeleteCampaign))

	// Product endpoints
	http.Handle("/get_all_products", auth.IsAuthorized(presentation.GetAllProducts))
	http.Handle("/get_product_info", auth.IsAuthorized(presentation.GetProductByProductCode))
	http.Handle("/create_product", auth.IsAuthorized(presentation.AddNewProduct))
	http.Handle("/update_product", auth.IsAuthorized(presentation.UpdateProduct))
	http.Handle("/delete_product", auth.IsAuthorized(presentation.DeleteProduct))

	// Order endpoints
	http.Handle("/get_all_orders", auth.IsAuthorized(presentation.GetAllOrder))
	http.Handle("/get_order_info", auth.IsAuthorized(presentation.GetOrderByProductCode))
	http.Handle("/create_order", auth.IsAuthorized(presentation.AddNewOrder))
	http.Handle("/update_order", auth.IsAuthorized(presentation.UpdateOrder))
	http.Handle("/delete_order", auth.IsAuthorized(presentation.DeleteOrder))

	http.ListenAndServe(":8080", nil)
}
