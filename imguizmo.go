package imguizmo

//go:generate go run ./generator

// #cgo windows LDFLAGS: -Ldist/windows
// #cgo linux LDFLAGS: -Ldist/linux
// #cgo darwin,amd64 LDFLAGS: -Ldist/macos/amd64
// #cgo darwin,arm64 LDFLAGS: -Ldist/macos/arm64
// #cgo LDFLAGS: -lcimguizmo -limguizmo
// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1
// #include "dist/include/cimguizmo.h"
import "C"
import (
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/nitrix/imgui-go"
)

// Operation enums.
const TranslateX = C.TRANSLATE_X
const TranslateY = C.TRANSLATE_Y
const TranslateZ = C.TRANSLATE_Z
const RotateX = C.ROTATE_X
const RotateY = C.ROTATE_Y
const RotateZ = C.ROTATE_Z
const RotateScreen = C.ROTATE_SCREEN
const ScaleX = C.SCALE_X
const ScaleY = C.SCALE_Y
const ScaleZ = C.SCALE_Z
const Bounds = C.BOUNDS
const ScaleXU = C.SCALE_XU
const ScaleYU = C.SCALE_YU
const ScaleZU = C.SCALE_ZU
const Translate = TranslateX | TranslateY | TranslateZ
const Rotate = RotateX | RotateY | RotateZ | RotateScreen
const Scale = ScaleX | ScaleY | ScaleZ
const ScaleU = ScaleXU | ScaleYU | ScaleZU
const Universal = Translate | Rotate | ScaleU

// Mode enums.
const Local = C.LOCAL
const World = C.WORLD

// Style enums.
const DirectionX = C.DIRECTION_X
const DirectionY = C.DIRECTION_Y
const DirectionZ = C.DIRECTION_Z
const PlaneX = C.PLANE_X
const PlaneY = C.PLANE_Y
const PlaneZ = C.PLANE_Z
const Selection = C.SELECTION
const Inactive = C.INACTIVE
const TranslationLine = C.TRANSLATION_LINE
const ScaleLine = C.SCALE_LINE
const RotationUsingBorder = C.ROTATION_USING_BORDER
const RotationUsingFill = C.ROTATION_USING_FILL
const HatchedAxisLines = C.HATCHED_AXIS_LINES
const Text = C.TEXT
const TextShadow = C.TEXT_SHADOW
const COUNT = C.COUNT

// Some types.
type Context = C.ImGuiContext
type Operation = C.OPERATION
type Mode = C.MODE

type Style struct {
	TranslationLineThickness   float32
	TranslationLineArrowSize   float32
	RotationLineThickness      float32
	RotationOuterLineThickness float32
	ScaleLineThickness         float32
	ScaleLineCircleSize        float32
	HatchedAxisLineThickness   float32
	CenterCircleSize           float32
	Colors                     [COUNT]mgl32.Vec4
}

func SetDrawlist(drawlist *imgui.DrawList) {
	C.ImGuizmo_SetDrawlist((*C.ImDrawList)(unsafe.Pointer(drawlist)))
}

func BeginFrame() {
	C.ImGuizmo_BeginFrame()
}

func SetImGuiContext(ctx *Context) {
	C.ImGuizmo_SetImGuiContext((*C.ImGuiContext)(ctx))
}

func IsOver_Nil() bool {
	return bool(C.ImGuizmo_IsOver_Nil())
}

func IsOver_OPERATION(op Operation) bool {
	return bool(C.ImGuizmo_IsOver_OPERATION(C.OPERATION(op)))
}

func IsUsing() bool {
	return bool(C.ImGuizmo_IsUsing())
}

func IsUsingAny() bool {
	return bool(C.ImGuizmo_IsUsingAny())
}

func Enable(enable bool) {
	C.ImGuizmo_Enable(C.bool(enable))
}

func SetRect(x, y, width, height float32) {
	C.ImGuizmo_SetRect(C.float(x), C.float(y), C.float(width), C.float(height))
}

func SetOrthographic(isOrthographic bool) {
	C.ImGuizmo_SetOrthographic(C.bool(isOrthographic))
}

func SetID(id int) {
	C.ImGuizmo_SetID(C.int(id))
}

func AllowAxisFlip(value bool) {
	C.ImGuizmo_AllowAxisFlip(C.bool(value))
}

func SetAxisLimit(value float32) {
	C.ImGuizmo_SetAxisLimit(C.float(value))
}

func SetPlaneLimit(value float32) {
	C.ImGuizmo_SetPlaneLimit(C.float(value))
}

func SetGizmoSizeClipSpace(value float32) {
	C.ImGuizmo_SetGizmoSizeClipSpace(C.float(value))
}

func DrawCubes(view mgl32.Mat4, projection mgl32.Mat4, matrices []mgl32.Mat4) {
	C.ImGuizmo_DrawCubes((*C.float)(&view[0]), (*C.float)(&projection[0]), (*C.float)(&matrices[0][0]), C.int(len(matrices)))
}

func DrawGrid(view mgl32.Mat4, projection mgl32.Mat4, matrix mgl32.Mat4, gridSize float32) {
	C.ImGuizmo_DrawGrid((*C.float)(&view[0]), (*C.float)(&projection[0]), (*C.float)(&matrix[0]), C.float(gridSize))
}

func Manipulate(view mgl32.Mat4, projection mgl32.Mat4, operation Operation, mode Mode, matrix *mgl32.Mat4, deltaMatrix *mgl32.Mat4) bool {
	return bool(C.ImGuizmo_Manipulate((*C.float)(&view[0]), (*C.float)(&projection[0]), C.OPERATION(operation), C.MODE(mode), (*C.float)(&matrix[0]), (*C.float)(&deltaMatrix[0]), nil, nil, nil))
}

func Manipulate_Snap(view mgl32.Mat4, projection mgl32.Mat4, operation Operation, mode Mode, matrix mgl32.Mat4, deltaMatrix mgl32.Mat4, snap mgl32.Vec3, localBounds mgl32.Vec3, boundsSnap mgl32.Vec3) bool {
	return bool(C.ImGuizmo_Manipulate((*C.float)(&view[0]), (*C.float)(&projection[0]), C.OPERATION(operation), C.MODE(mode), (*C.float)(&matrix[0]), (*C.float)(&deltaMatrix[0]), (*C.float)(&snap[0]), (*C.float)(&localBounds[0]), (*C.float)(&boundsSnap[0])))
}

func DecomposeMatrixToComponents(matrix mgl32.Mat4) (mgl32.Vec3, mgl32.Vec3, mgl32.Vec3) {
	var translation, rotation, scale mgl32.Vec3
	C.ImGuizmo_DecomposeMatrixToComponents((*C.float)(&matrix[0]), (*C.float)(&translation[0]), (*C.float)(&rotation[0]), (*C.float)(&scale[0]))
	return translation, rotation, scale
}

func RecomposeMatrixFromComponents(translation, rotation, scale mgl32.Vec3) mgl32.Mat4 {
	var matrix mgl32.Mat4
	C.ImGuizmo_RecomposeMatrixFromComponents((*C.float)(&translation[0]), (*C.float)(&rotation[0]), (*C.float)(&scale[0]), (*C.float)(&matrix[0]))
	return matrix
}

func ViewManipulate(view *mgl32.Mat4, length float32, position, size mgl32.Vec2, backgroundColor uint32) {
	vposition := C.ImVec2{C.float(position[0]), C.float(position[1])}
	vsize := C.ImVec2{C.float(size[0]), C.float(size[1])}
	C.ImGuizmo_ViewManipulate_Float((*C.float)(&view[0]), C.float(length), vposition, vsize, C.ImU32(backgroundColor))
}

func ViewManipulate_FloatPtr(view, projection *mgl32.Mat4, operation Operation, mode Mode, matrix *mgl32.Mat4, length float32, position, size mgl32.Vec2, backgroundColor uint32) {
	vposition := C.ImVec2{C.float(position[0]), C.float(position[1])}
	vsize := C.ImVec2{C.float(size[0]), C.float(size[1])}
	C.ImGuizmo_ViewManipulate_FloatPtr((*C.float)(&view[0]), (*C.float)(&projection[0]), C.OPERATION(operation), C.MODE(mode), (*C.float)(&matrix[0]), C.float(length), vposition, vsize, C.ImU32(backgroundColor))
}

func Style_Style() *Style {
	cStyle := C.Style_Style()
	goStyle := (*Style)(unsafe.Pointer(cStyle))
	return goStyle
}

func Style_destroy(self *Style) {
	C.Style_destroy((*C.Style)(unsafe.Pointer(self)))
}

func GetStyle() *Style {
	cStyle := C.ImGuizmo_GetStyle()
	goStyle := (*Style)(unsafe.Pointer(cStyle))
	return goStyle
}
