# Lalamove Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/jamie0xgitc0decat/lalamove-go-sdk.svg)](https://pkg.go.dev/github.com/jamie0xgitc0decat/lalamove-go-sdk)

**Lalamove Go SDK** is a Go library that provides seamless integration with Lalamove’s Official API V3. Our aim is to deliver full compatibility with Lalamove’s API, enabling developers to easily incorporate a broad range of logistics and order management features into their Go applications.

---

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
  - [Search](#search)
  - [Quotation Details](#quotation-details)
    - [Quotation ID](#quotation-id)
    - [Quotation Validity Period](#quotation-validity-period)
    - [Price Breakdown](#price-breakdown)
    - [Distance](#distance)
  - [Change Driver](#change-driver)
  - [Route Optimization](#route-optimization)
  - [Authentication](#authentication)
  - [Order Flow](#order-flow)
    - [Available Markets](#available-markets)
    - [Service Types & Special Requests](#service-types--special-requests)
    - [Get Quotation](#get-quotation)
    - [Quotation Details](#quotation-details-1)
    - [Place Order](#place-order)
    - [Order Details](#order-details)
    - [Driver Details](#driver-details)
    - [Add Priority Fee](#add-priority-fee)
    - [Edit Order](#edit-order)
    - [Cancel Order](#cancel-order)
    - [Change Driver (Order)](#change-driver-order)
  - [Get City Info](#get-city-info)
  - [Webhook](#webhook)
  - [Errors](#errors)
- [Getting on Board](#getting-on-board)
- [Tutorial](#tutorial)
- [Support](#support)
- [Roadmap](#roadmap)
- [Change Log](#change-log)
- [License](#license)

---

## Introduction

The **Lalamove Go SDK** is a robust and comprehensive client library built in Go for interacting with Lalamove’s Official API V3. This SDK aims to cover the full spectrum of Lalamove functionalities, ensuring you can easily integrate searching, quoting, ordering, and managing deliveries directly within your Go applications.

---

## Features

### Search

Perform searches across available services, routes, and drivers with ease. Customize your query parameters to find the exact information you need.

### Quotation Details

- **Quotation ID:**  
  Each quotation is uniquely identified to facilitate tracking and reference.

- **Quotation Validity Period:**  
  The validity period during which the quotation remains active. Renew or request a new quotation once expired.

- **Price Breakdown:**  
  Detailed cost components including base rates, fees, taxes, and additional charges that explain the final price.

- **Distance:**  
  Calculates the distance between pickup and drop-off points, influencing the quotation and route planning.

### Change Driver

Request or assign a different driver for an order. This endpoint allows you to manage driver changes seamlessly.

### Route Optimization

Retrieve the most efficient route based on real-time conditions, traffic data, and other dynamic factors to reduce delivery time.

### Authentication

Secure your API interactions with token-based authentication. Generate and manage your API tokens to authenticate each request.

### Order Flow

Manage the entire lifecycle of an order with a series of endpoints:

- **Available Markets:**  
  Retrieve a list of markets where Lalamove services are available.

- **Service Types & Special Requests:**  
  Discover and select the service types offered, and submit any special requests required for your order.

- **Get Quotation:**  
  Request a quotation by supplying necessary details such as locations, service type, and other parameters.

- **Quotation Details:**  
  Access detailed information on the provided quotation, including pricing and validity.

- **Place Order:**  
  Confirm and submit an order based on the received quotation.

- **Order Details:**  
  Retrieve comprehensive details about your placed order, including current status and driver information.

- **Driver Details:**  
  Get contact and vehicle information about the driver assigned to your order.

- **Add Priority Fee:**  
  Boost your order’s priority by adding an extra fee, ensuring faster processing.

- **Edit Order:**  
  Modify your order details such as pickup time or location after placement.

- **Cancel Order:**  
  Cancel an order that hasn’t been fulfilled, subject to cancellation policies.

- **Change Driver (Order):**  
  Reassign a different driver if needed after the order is in progress.

### Get City Info

Retrieve detailed operational information about a city, including available services, local regulations, and operational details.

### Webhook

Set up webhooks to receive real-time notifications on order status updates, driver changes, and other important events.

### Errors

A detailed guide on error codes and messages returned by the API, along with troubleshooting tips and solutions.

---

## Getting on Board

A quick-start guide to help new users integrate the SDK:

1. **Prerequisites:**  
   Ensure you have Go installed and a valid Lalamove API token.

2. **Installation:**  
   Install the SDK using:
   ```bash
   go get github.com/jamie0xgitc0decat/lalamove-go-sdk
