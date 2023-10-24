// Common data structures
// Not used

package data

import "time"

// An example payload to send between our server and client.
type BodyMetric struct {
	T       time.Time `json:"time"`
	Hr      float32   `json:"heart_rate"`
	Speed   float32   `json:"speed"`
	Cadence float32   `json:"cadence"`
	Temp    int32     `json:"temperature"`
}
