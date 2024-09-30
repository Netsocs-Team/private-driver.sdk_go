package objects

type ObjectChild interface {
	IntoObject() Object
}

type Object struct {
	ObjectID        ObjectID        `json:"object_id"`        // Unique identifier for the object
	Name            string          `json:"name"`             // Name of the object
	Type            ObjectType      `json:"type"`             // Type of the object
	State           int             `json:"state"`            // JSON string of the state of the object
	StatesAvailable []StateObject   `json:"states_available"` // JSON string of the available states of the object
	StateAttributes StateAttributes `json:"state_attributes"` // JSON string (only key value) of the state attributes of the object
	DeviceID        int             `json:"device_id"`        // Rel to device in devices netsocs table
	Enabled         bool            `json:"enabled"`          // Is the object enabled
	Icon            string          `json:"icon"`             // Icon for the object. Default Netsocs use FontAwesome, but can be custom
}

type ObjectID string

type StateObject struct {
	StateID string `json:"state_id"` // Unique identifier for the state
	// Name of the state. This is used to display the state in the frontend and should be unique for each object.
	// For example. map["en"] = "On", map["de"] = "An", map["fr"] = "Sur", map["es"] = "Encendido"
	InternationalizedName map[string]string `json:"name"`
	Icon                  string            `json:"icon"` // Icon for the state. Default Netsocs use FontAwesome, but can be custom
}

const NullState = -1

type StateAttributes map[string]string

type ObjectType string
