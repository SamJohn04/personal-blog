package model

import "time"

type BlogTitle struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

type BlogPost struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}
