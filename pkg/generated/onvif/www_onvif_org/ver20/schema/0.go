// Code generated by gowsdl DO NOT EDIT.

package schema

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"

	tt "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"

	bd "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/analytics/humanbody"

	fc "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/analytics/humanface"

	wsnt "github.com/eyetowers/gonvif/pkg/generated/onvif/docs_oasisopen_org/wsn/b2"
)

// against "unused imports"
var _ time.Time
var _ xml.Name
var _ context.Context
var _ soap.SOAPEnvelope

type VehicleType string

const (
	VehicleTypeBus VehicleType = "Bus"

	VehicleTypeCar VehicleType = "Car"

	VehicleTypeTruck VehicleType = "Truck"

	VehicleTypeBicycle VehicleType = "Bicycle"

	VehicleTypeMotorcycle VehicleType = "Motorcycle"
)

type PlateType string

const (
	PlateTypeNormal PlateType = "Normal"

	PlateTypePolice PlateType = "Police"

	PlateTypeDiplomat PlateType = "Diplomat"

	PlateTypeTemporary PlateType = "Temporary"
)

type ObjectType string

const (
	ObjectTypeAnimal ObjectType = "Animal"

	ObjectTypeHumanFace ObjectType = "HumanFace"

	ObjectTypeHuman ObjectType = "Human"

	ObjectTypeBicycle ObjectType = "Bicycle"

	ObjectTypeVehicle ObjectType = "Vehicle"

	ObjectTypeLicensePlate ObjectType = "LicensePlate"

	ObjectTypeBike ObjectType = "Bike"
)

type ClassType string

const (
	ClassTypeAnimal ClassType = "Animal"

	ClassTypeFace ClassType = "Face"

	ClassTypeHuman ClassType = "Human"

	ClassTypeVehical ClassType = "Vehical"

	ClassTypeOther ClassType = "Other"
)

type Appearance struct {
	Transformation *tt.Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty"`

	Shape *ShapeDescriptor `xml:"Shape,omitempty" json:"Shape,omitempty"`

	Color *tt.ColorDescriptor `xml:"Color,omitempty" json:"Color,omitempty"`

	Class *ClassDescriptor `xml:"Class,omitempty" json:"Class,omitempty"`

	Extension *AppearanceExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	GeoLocation *tt.GeoLocation `xml:"GeoLocation,omitempty" json:"GeoLocation,omitempty"`

	VehicleInfo []*VehicleInfo `xml:"VehicleInfo,omitempty" json:"VehicleInfo,omitempty"`

	LicensePlateInfo *LicensePlateInfo `xml:"LicensePlateInfo,omitempty" json:"LicensePlateInfo,omitempty"`

	HumanFace *fc.HumanFace `xml:"HumanFace,omitempty" json:"HumanFace,omitempty"`

	HumanBody *bd.HumanBody `xml:"HumanBody,omitempty" json:"HumanBody,omitempty"`

	ImageRef string `xml:"ImageRef,omitempty" json:"ImageRef,omitempty"`

	Image []byte `xml:"Image,omitempty" json:"Image,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AppearanceExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VehicleInfo struct {
	Type *StringLikelihood `xml:"Type,omitempty" json:"Type,omitempty"`

	Brand *StringLikelihood `xml:"Brand,omitempty" json:"Brand,omitempty"`

	Model *StringLikelihood `xml:"Model,omitempty" json:"Model,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type LicensePlateInfo struct {

	// A string of vehicle license plate number.
	PlateNumber *StringLikelihood `xml:"PlateNumber,omitempty" json:"PlateNumber,omitempty"`

	// A description of the vehicle license plate, e.g., "Normal", "Police", "Diplomat"
	PlateType *StringLikelihood `xml:"PlateType,omitempty" json:"PlateType,omitempty"`

	// Describe the country of the license plate, in order to avoid the same license plate number.
	CountryCode *StringLikelihood `xml:"CountryCode,omitempty" json:"CountryCode,omitempty"`

	// State province or authority that issue the license plate.
	IssuingEntity *StringLikelihood `xml:"IssuingEntity,omitempty" json:"IssuingEntity,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ShapeDescriptor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Shape" json:"-"`

	BoundingBox *tt.Rectangle `xml:"BoundingBox,omitempty" json:"BoundingBox,omitempty"`

	CenterOfGravity *tt.Vector `xml:"CenterOfGravity,omitempty" json:"CenterOfGravity,omitempty"`

	Polygon []*tt.Polygon `xml:"Polygon,omitempty" json:"Polygon,omitempty"`

	Extension *ShapeDescriptorExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ShapeDescriptorExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type StringLikelihood struct {
	Value string `xml:",chardata" json:",omitempty"`

	Likelihood float32 `xml:"http://www.onvif.org/ver10/schema Likelihood,attr,omitempty" json:"Likelihood,omitempty"`
}

type ClassDescriptor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Class" json:"-"`

	ClassCandidate []struct {
		Type *ClassType `xml:"Type,omitempty" json:"Type,omitempty"`

		Likelihood float32 `xml:"Likelihood,omitempty" json:"Likelihood,omitempty"`
	} `xml:"ClassCandidate,omitempty" json:"ClassCandidate,omitempty"`

	Extension *ClassDescriptorExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// ONVIF recommends to use this 'Type' element instead of 'ClassCandidate' and 'Extension' above for new design. Acceptable values are defined in tt:ObjectType.
	Type []*StringLikelihood `xml:"Type,omitempty" json:"Type,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ClassDescriptorExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	OtherTypes []*OtherType `xml:"OtherTypes,omitempty" json:"OtherTypes,omitempty"`

	Extension *ClassDescriptorExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ClassDescriptorExtension2 struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type OtherType struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema OtherTypes" json:"-"`

	// Object Class Type
	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	// A likelihood/probability that the corresponding object belongs to this class. The sum of the likelihoods shall NOT exceed 1
	Likelihood float32 `xml:"Likelihood,omitempty" json:"Likelihood,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Object struct {
	*ObjectId

	Appearance *Appearance `xml:"Appearance,omitempty" json:"Appearance,omitempty"`

	Behaviour *Behaviour `xml:"Behaviour,omitempty" json:"Behaviour,omitempty"`

	Extension *ObjectExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Object ID of the parent object. eg: License plate object has Vehicle object as parent.

	Parent int32 `xml:"http://www.onvif.org/ver10/schema Parent,attr,omitempty" json:"Parent,omitempty"`
}

type ObjectExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Frame struct {
	PTZStatus *tt.PTZStatus `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty"`

	Transformation *tt.Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty"`

	Object []*Object `xml:"Object,omitempty" json:"Object,omitempty"`

	ObjectTree *ObjectTree `xml:"ObjectTree,omitempty" json:"ObjectTree,omitempty"`

	Extension *FrameExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	SceneImageRef string `xml:"SceneImageRef,omitempty" json:"SceneImageRef,omitempty"`

	SceneImage []byte `xml:"SceneImage,omitempty" json:"SceneImage,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	UtcTime soap.XSDDateTime `xml:"http://www.onvif.org/ver10/schema UtcTime,attr,omitempty" json:"UtcTime,omitempty"`

	// Default color space of Color definitions in frame. Valid values are "RGB" and "YCbCr". Defaults to "YCbCr".

	Colorspace string `xml:"http://www.onvif.org/ver10/schema Colorspace,attr,omitempty" json:"Colorspace,omitempty"`

	// Optional name of the analytics module that generated this frame.

	Source string `xml:"http://www.onvif.org/ver10/schema Source,attr,omitempty" json:"Source,omitempty"`
}

type FrameExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	MotionInCells *MotionInCells `xml:"MotionInCells,omitempty" json:"MotionInCells,omitempty"`

	Extension *FrameExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FrameExtension2 struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Merge struct {
	From []*ObjectId `xml:"from,omitempty" json:"from,omitempty"`

	To *ObjectId `xml:"to,omitempty" json:"to,omitempty"`
}

type Split struct {
	From *ObjectId `xml:"from,omitempty" json:"from,omitempty"`

	To []*ObjectId `xml:"to,omitempty" json:"to,omitempty"`
}

type Rename struct {
	From *ObjectId `xml:"from,omitempty" json:"from,omitempty"`

	To *ObjectId `xml:"to,omitempty" json:"to,omitempty"`
}

type ObjectId struct {
	ObjectId int32 `xml:"http://www.onvif.org/ver10/schema ObjectId,attr,omitempty" json:"ObjectId,omitempty"`
}

type Behaviour struct {
	Removed struct {
	} `xml:"Removed,omitempty" json:"Removed,omitempty"`

	Idle struct {
	} `xml:"Idle,omitempty" json:"Idle,omitempty"`

	Extension *BehaviourExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Speed float32 `xml:"Speed,omitempty" json:"Speed,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type BehaviourExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ObjectTree struct {
	Rename []*Rename `xml:"Rename,omitempty" json:"Rename,omitempty"`

	Split []*Split `xml:"Split,omitempty" json:"Split,omitempty"`

	Merge []*Merge `xml:"Merge,omitempty" json:"Merge,omitempty"`

	Delete []*ObjectId `xml:"Delete,omitempty" json:"Delete,omitempty"`

	Extension *ObjectTreeExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ObjectTreeExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MotionInCells struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Number of columns of the cell grid (x dimension)

	Columns int32 `xml:"http://www.onvif.org/ver10/schema Columns,attr,omitempty" json:"Columns,omitempty"`

	// Number of rows of the cell grid (y dimension)

	Rows int32 `xml:"http://www.onvif.org/ver10/schema Rows,attr,omitempty" json:"Rows,omitempty"`

	// A “1” denotes a cell where motion is detected and a “0” an empty cell. The first cell is in the upper left corner. Then the cell order goes first from left to right and then from up to down.  If the number of cells is not a multiple of 8 the last byte is filled with zeros. The information is run length encoded according to Packbit coding in ISO 12369 (TIFF, Revision 6.0).

	Cells []byte `xml:"http://www.onvif.org/ver10/schema Cells,attr,omitempty" json:"Cells,omitempty"`
}

type MetadataStream struct {
	VideoAnalytics *VideoAnalyticsStream `xml:"VideoAnalytics,omitempty" json:"VideoAnalytics,omitempty"`

	PTZ *PTZStream `xml:"PTZ,omitempty" json:"PTZ,omitempty"`

	Event *EventStream `xml:"Event,omitempty" json:"Event,omitempty"`

	Extension *MetadataStreamExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type MetadataStreamExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	AudioAnalyticsStream *AudioAnalyticsStream `xml:"AudioAnalyticsStream,omitempty" json:"AudioAnalyticsStream,omitempty"`

	Extension *MetadataStreamExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataStreamExtension2 struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioAnalyticsStream struct {
	AudioDescriptor []*AudioDescriptor `xml:"AudioDescriptor,omitempty" json:"AudioDescriptor,omitempty"`

	Extension *AudioAnalyticsStreamExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioDescriptor struct {
	Items []string `xml:",any" json:"items,omitempty"`

	UtcTime soap.XSDDateTime `xml:"http://www.onvif.org/ver10/schema UtcTime,attr,omitempty" json:"UtcTime,omitempty"`
}

type AudioAnalyticsStreamExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoAnalyticsStream struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema VideoAnalytics" json:"-"`

	Frame *Frame `xml:"Frame,omitempty" json:"Frame,omitempty"`

	Extension *VideoAnalyticsStreamExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type VideoAnalyticsStreamExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZStream struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema PTZ" json:"-"`

	PTZStatus *tt.PTZStatus `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty"`

	Extension *PTZStreamExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZStreamExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type EventStream struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Event" json:"-"`

	NotificationMessage *wsnt.NotificationMessage `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty"`

	Extension *EventStreamExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type EventStreamExtension struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Extension" json:"-"`

	Items []string `xml:",any" json:"items,omitempty"`
}
