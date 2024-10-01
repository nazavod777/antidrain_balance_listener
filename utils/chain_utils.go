package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

func GetNonce(client *fasthttp.Client, rpcUrl string, address string) (uint64, error) {
	for {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI(rpcUrl)
		req.Header.SetMethod("POST")
		req.Header.Set("Content-Type", "application/json")

		params := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_getTransactionCount",
			"params": []interface{}{
				address,
				"pending",
			},
			"id": 1,
		}
		body, err := json.Marshal(params)
		if err != nil {
			log.Printf("Error When Marhsalling JSON When Getting Nonce: %v", err)
		}
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err := client.Do(req, resp); err != nil {
			log.Printf("Error When Making Request When Getting Nonce: %v", err)
			continue
		}

		var responseJson map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &responseJson); err != nil {
			log.Printf("Error When Decoding Response When Getting Nonce: %v", err)
			continue
		}

		if result, ok := responseJson["result"].(string); ok {
			nonce, err := strconv.ParseInt(result, 0, 64)
			if err != nil {
				log.Printf("Error When Converting Nonce When Getting Nonce: %v", err)
				continue
			}
			return uint64(nonce), nil
		}

		log.Printf("Error In Json Response When Getting Nonce: %s", resp.Body())
	}
}

func GetBalance(client *fasthttp.Client, rpcUrl string, address string) (int, error) {
	for {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI(rpcUrl)
		req.Header.SetMethod("POST")
		req.Header.Set("Content-Type", "application/json")

		params := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_getBalance",
			"params": []interface{}{
				address,
				"pending",
			},
			"id": 1,
		}
		body, err := json.Marshal(params)
		if err != nil {
			log.Printf("Error When Marhsalling JSON When Getting Balance: %v", err)
		}
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err := client.Do(req, resp); err != nil {
			log.Printf("Error When Making Request When Getting Balance: %v", err)
			continue
		}

		var responseJson map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &responseJson); err != nil {
			log.Printf("Error When Decoding Response When Getting Balance: %v", err)
			continue
		}

		if result, ok := responseJson["result"].(string); ok {
			balance, err := strconv.ParseInt(result, 0, 64)
			if err != nil {
				log.Printf("Error When Converting Balance When Getting Balance: %v", err)
				continue
			}
			return int(balance), nil
		}

		log.Printf("Error In Json Response When Getting Balance: %s", resp.Body())
	}
}

func GetGwei(client *fasthttp.Client, rpcUrl string) (int, error) {
	for {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI(rpcUrl)
		req.Header.SetMethod("POST")
		req.Header.Set("Content-Type", "application/json")

		params := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_gasPrice",
			"params":  []interface{}{},
			"id":      1,
		}
		body, err := json.Marshal(params)
		if err != nil {
			log.Fatalf("Error When Marhsalling JSON When Getting GWEI: %v", err)
		}
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err := client.Do(req, resp); err != nil {
			log.Printf("Error When Making Request When Getting GWEI: %v", err)
			continue
		}

		var responseJson map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &responseJson); err != nil {
			log.Printf("Error When Decoding Response When Getting GWEI: %v", err)
			continue
		}

		if result, ok := responseJson["result"].(string); ok {
			nonce, err := strconv.ParseInt(result, 0, 64)
			if err != nil {
				log.Printf("Error When Converting GWEI When Getting GWEI: %v", err)
				continue
			}
			return int(nonce), nil
		}

		log.Printf("Error In Json Response When Getting GWEI: %s", resp.Body())
	}
}

func GetChainID(client *fasthttp.Client, rpcUrl string) (uint64, error) {
	for {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI(rpcUrl)
		req.Header.SetMethod("POST")
		req.Header.Set("Content-Type", "application/json")

		params := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_chainId",
			"params":  []interface{}{},
			"id":      1,
		}
		body, err := json.Marshal(params)
		if err != nil {
			log.Fatalf("Error When Marhsalling JSON When Getting ChainID: %v", err)
		}
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if err := client.Do(req, resp); err != nil {
			log.Printf("Error When Making Request When Getting ChainID: %v", err)
			continue
		}

		var responseJson map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &responseJson); err != nil {
			log.Printf("Error When Decoding Response When Getting ChainID: %v", err)
			continue
		}

		if result, ok := responseJson["result"].(string); ok {
			nonce, err := strconv.ParseInt(result, 0, 64)
			if err != nil {
				log.Printf("Error When Converting ChainID When Getting ChainID: %v", err)
				continue
			}
			return uint64(nonce), nil
		}

		log.Printf("Error In Json Response When Getting ChainID: %s", resp.Body())
	}
}

func SendTransaction(client *fasthttp.Client, rpcUrl string, rawTransactionData string) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(rpcUrl)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_sendRawTransaction",
		"params": []interface{}{
			rawTransactionData,
		},
		"id": 1,
	}
	body, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Error When Marhsalling JSON When Sending Transaction: %v", err)
	}
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		log.Printf("Error When Making Request When Sending Transaction: %v", err)
		return "", err
	}

	var responseJson map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseJson); err != nil {
		log.Printf("Error When Decoding Response When Sending Transaction: %v", err)
		return "", err
	}

	if result, ok := responseJson["result"].(string); ok {
		return result, nil
	}

	log.Printf("Error In Json Response When Sending Transaction: %s", resp.Body())
	return "", nil
}
