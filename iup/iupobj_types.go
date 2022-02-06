package iup

import (
	"unsafe"
)

// Iup callback IUP_DEFAULT_ACTION
type CommonDefaultAction struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_MAP_CB
type CommonMap struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_UNMAP_CB
type CommonUnmap struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_DESTROY_CB
type CommonDestroy struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_GETFOCUS_CB
type CommonGetFocus struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_KILLFOCUS_CB
type CommonKillFocus struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_ENTERWINDOW_CB
type CommonEnterWindow struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_LEAVEWINDOW_CB
type CommonLeaveWindow struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_HELP_CB
type CommonHelp struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_K_ANY
type CommonKeyAny struct {
	Sender *Handle
	Return int32
	Key    int32
}

// Iup callback IUP_CLOSE_CB
type DialogClose struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_RESIZE_CB
type DialogResize struct {
	Sender *Handle
	Return int32
	Width  int32
	Height int32
}

// Iup callback IUP_DROPFILES_CB
type DialogDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_SHOW_CB
type DialogShow struct {
	Sender *Handle
	Return int32
	State  int32
}

// Iup callback IUP_FILE_CB
type FileDlgFile struct {
	Sender   *Handle
	Return   int32
	FileName string
	Status   string
}

// Iup callback IUP_ACTION
type ItemAction struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_HIGHLIGHT_CB
type ItemHighlight struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_OPEN_CB
type MenuOpen struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_MENUCLOSE_CB
type MenuClose struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_HIGHLIGHT_CB
type SubMenuHighlight struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_ACTION
type TimerAction struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_ACTION
type ButtonAction struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_BUTTON_CB
type MouseButton struct {
	Sender  *Handle
	Return  int32
	Button  int32
	Pressed int32
	X       int32
	Y       int32
	Status  string
}

// Iup callback IUP_ACTION
type CanvasAction struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_DROPFILES_CB
type CanvasDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_FOCUS_CB
type CanvasFocus struct {
	Sender *Handle
	Return int32
	Focus  int32
}

// Iup callback IUP_MOTION_CB
type MouseMotion struct {
	Sender *Handle
	Return int32
	X      int32
	Y      int32
	Status string
}

//PLOTMOTION_CB
type PlotMotion struct {
	Sender *Handle
	Return int32
	X      float64
	Y      float64
	Status string
}

//XTICKFORMATNUMBER_CB
type PlotXtick struct {
	Sender         *Handle
	Return         int32
	Buffer, Format string
	X              float64
	Decimal_symbol string
}

// Iup callback IUP_KEYPRESS_CB
type CanvasKeyPress struct {
	Sender *Handle
	Return int32
	C      int32
	Press  int32
}

// Iup callback IUP_RESIZE_CB
type CanvasResize struct {
	Sender *Handle
	Return int32
	Width  int32
	Height int32
}

// Iup callback IUP_SCROLL_CB
type CanvasScroll struct {
	Sender *Handle
	Return int32
	Op     int32
	PosX   float32
	PosY   float32
}

// Iup callback IUP_TOUCH_CB
type TouchEvent struct {
	Sender *Handle
	Return int32
	X      int32
	Y      int32
	State  string
}

// Iup callback IUP_MULTITOUCH_CB
type MultiTouchEvent struct {
	Sender *Handle
	Return int32
	Count  int32
	Pid    *int32
	Px     *int32
	Py     *int32
	PState *int32
}

// Iup callback IUP_WHEEL_CB
type CanvasWheel struct {
	Sender *Handle
	Return int32
	Delta  float32
	X      int32
	Y      int32
	Status string
}

// Iup callback IUP_WOM_CB
type CanvasWom struct {
	Sender *Handle
	Return int32
	State  int32
}

// Iup callback IUP_DROPFILES_CB
type LabelDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_ACTION
type ListAction struct {
	Sender *Handle
	Return int32
	Text   string
	Item   int32
	State  int32
}

// Iup callback IUP_CARET_CB
type ListCaret struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Pos    int32
}

// Iup callback IUP_DBLCLICK_CB
type ListDblclick struct {
	Sender *Handle
	Return int32
	Item   int32
	Text   string
}

// Iup callback IUP_DROPDOWN_CB
type ListDropDown struct {
	Sender *Handle
	Return int32
	State  int32
}

// Iup callback IUP_DROPFILES_CB
type ListDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_EDIT_CB
type ListEdit struct {
	Sender   *Handle
	Return   int32
	C        int32
	NewValue string
}

// Iup callback IUP_MULTISELECT_CB
type ListMultiSelect struct {
	Sender *Handle
	Return int32
	Value  string
}

// Iup callback IUP_VALUECHANGED_CB
type ValueChanged struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_SPIN_CB
type SpinInc struct {
	Sender *Handle
	Return int32
	Inc    int32
}

// Iup callback IUP_SPIN_CB
type SpinBoxInc struct {
	Sender *Handle
	Return int32
	Inc    int32
}

// Iup callback IUP_TABCHANGE_CB
type TabsChange struct {
	Sender *Handle
	Return int32
	NewTab *Handle
	OldTab *Handle
}

// Iup callback IUP_TABCHANGEPOS_CB
type TabsChangePos struct {
	Sender *Handle
	Return int32
	NewPos int32
	OldPos int32
}

// Iup callback IUP_ACTION
type TextAction struct {
	Sender   *Handle
	Return   int32
	C        int32
	NewValue string
}

// Iup callback IUP_CARET_CB
type TextCaret struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Pos    int32
}

// Iup callback IUP_DROPFILES_CB
type TextDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_SPIN_CB
type TextSpin struct {
	Sender *Handle
	Return int32
	Inc    int32
}

// Iup callback IUP_ACTION
type ToggleAction struct {
	Sender *Handle
	Return int32
	State  int32
}

// Iup callback IUP_SELECTION_CB
type TreeSelection struct {
	Sender *Handle
	Return int32
	Id     int32
	Status int32
}

// Iup callback IUP_MULTISELECTION_CB
type TreeMultiSelection struct {
	Sender *Handle
	Return int32
	Ids    *int32
	N      int32
}

// Iup callback IUP_BRANCHOPEN_CB
type TreeBranchOpen struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_BRANCHCLOSE_CB
type TreeBranchClose struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_EXECUTELEAF_CB
type TreeExecuteLeaf struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_EXECUTEBRANCH_CB
type TreeExecuteBranch struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_SHOWRENAME_CB
type TreeShowRename struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_RENAME_CB
type TreeRename struct {
	Sender *Handle
	Return int32
	Id     int32
	Title  string
}

// Iup callback IUP_DRAGDROP_CB
type TreeDragDrop struct {
	Sender    *Handle
	Return    int32
	DragId    int32
	DropId    int32
	IsShift   int32
	IsControl int32
}

// Iup callback IUP_NODEREMOVED_CB
type TreeNodeRemoved struct {
	Sender   *Handle
	Return   int32
	UserData unsafe.Pointer
}

// Iup callback IUP_RIGHTCLICK_CB
type TreeRightClick struct {
	Sender *Handle
	Return int32
	Id     int32
}

// Iup callback IUP_DROPFILES_CB
type TreeDropFiles struct {
	Sender   *Handle
	Return   int32
	FileName string
	Num      int32
	X        int32
	Y        int32
}

// Iup callback IUP_DRAW_CB
type CellsDraw struct {
	Sender *Handle
	Return int32
	Line   int32
	Column int32
	Xmin   int32
	Xmax   int32
	Ymin   int32
	Ymax   int32
	Cnv    unsafe.Pointer
}

// Iup callback IUP_HEIGHT_CB
type CellsHeight struct {
	Sender *Handle
	Return int32
	Line   int32
}

// Iup callback IUP_HSPAN_CB
type CellsHspan struct {
	Sender *Handle
	Return int32
	Line   int32
	Column int32
}

// Iup callback IUP_MOUSECLICK_CB
type CellsMouseClick struct {
	Sender  *Handle
	Return  int32
	Button  int32
	Pressed int32
	Line    int32
	Column  int32
	X       int32
	Y       int32
	Status  string
}

// Iup callback IUP_MOUSEMOTION_CB
type CellsMouseMotion struct {
	Sender *Handle
	Return int32
	Line   int32
	Column int32
	X      int32
	Y      int32
	R      string
}

// Iup callback IUP_NCOLS_CB
type CellsNcols struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_NLINES_CB
type CellsNlines struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_SCROLLING_CB
type CellsScrolling struct {
	Sender *Handle
	Return int32
	Line   int32
	Column int32
}

// Iup callback IUP_VSPAN_CB
type CellsVspan struct {
	Sender *Handle
	Return int32
	Line   int32
	Column int32
}

// Iup callback IUP_WIDTH_CB
type CellsWidth struct {
	Sender *Handle
	Return int32
	Column int32
}

// Iup callback IUP_CELL_CB
type ColorbarCell struct {
	Sender *Handle
	Return string
	Cell   int32
}

// Iup callback IUP_EXTENDED_CB
type ColorbarExtended struct {
	Sender *Handle
	Return int32
	Cell   int32
}

// Iup callback IUP_SELECT_CB
type ColorbarSelect struct {
	Sender *Handle
	Return int32
	Cell   int32
	Type   int32
}

// Iup callback IUP_SWITCH_CB
type ColorbarSwitch struct {
	Sender   *Handle
	Return   int32
	PrimCell int32
	SecCell  int32
}

// Iup callback IUP_CHANGE_CB
type ColorBrowserChange struct {
	Sender *Handle
	Return int32
	R      byte
	G      byte
	B      byte
}

// Iup callback IUP_DRAG_CB
type ColorBrowserDrag struct {
	Sender *Handle
	Return int32
	R      byte
	G      byte
	B      byte
}

// Iup callback IUP_BUTTON_PRESS_CB
type DialButtonPress struct {
	Sender *Handle
	Return int32
	Angle  float64
}

// Iup callback IUP_BUTTON_RELEASE_CB
type DialButtonRelease struct {
	Sender *Handle
	Return int32
	Angle  float64
}

// Iup callback IUP_MOUSEMOVE_CB
type DialMouseMove struct {
	Sender *Handle
	Return int32
	Angle  float64
}

// Iup callback IUP_ACTION_CB
type MatrixAction struct {
	Sender  *Handle
	Return  int32
	Key     int32
	Lin     int32
	Col     int32
	Edition int32
	Value   string
}

// Iup callback IUP_CLICK_CB
type MatrixClick struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Status string
}

// Iup callback IUP_RELEASE_CB
type MatrixRelease struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Status string
}

// Iup callback IUP_MOUSEMOVE_CB
type MatrixMouseMove struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_ENTERITEM_CB
type MatrixEnterItem struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_LEAVEITEM_CB
type MatrixLeaveItem struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_SCROLLTOP_CB
type MatrixScrollTop struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_BGCOLOR_CB
type MatrixBgcolor struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Red    *uint
	Green  *uint
	Blue   *uint
}

// Iup callback IUP_FGCOLOR_CB
type MatrixFgcolor struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Red    *uint
	Green  *uint
	Blue   *uint
}

// Iup callback IUP_FONT_CB
type MatrixFont struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_DRAW_CB
type MatrixDraw struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	X1     int32
	X2     int32
	Y1     int32
	Y2     int32
	Cnv    unsafe.Pointer
}

// Iup callback IUP_DROPCHECK_CB
type MatrixDropCheck struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_DROP_CB
type MatrixDrop struct {
	Sender *Handle
	Drop   *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_DROPSELECT_CB
type MatrixDropSelect struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Drop   *Handle
	T      string
	I      int32
	V      int32
}

// Iup callback IUP_EDITION_CB
type MatrixEdition struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Mode   int32
	Update int32
}

// Iup callback IUP_VALUE_CB
type MatrixValue struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_VALUE_EDIT_CB
type MatrixValueEdit struct {
	Sender   *Handle
	Return   int32
	Lin      int32
	Col      int32
	NewValue string
}

// Iup callback IUP_MARK_CB
type MatrixMark struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
}

// Iup callback IUP_MARKEDIT_CB
type MatrixMarkEdit struct {
	Sender *Handle
	Return int32
	Lin    int32
	Col    int32
	Marked int32
}

// Iup callback IUP_DELETE_CB
type PPlotDelete struct {
	Sender      *Handle
	Return      int32
	Index       int32
	SampleIndex int32
	X           float32
	Y           float32
}

// Iup callback IUP_DELETEBEGIN_CB
type PPlotDeleteBegin struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_DELETEEND_CB
type PPlotDeleteEnd struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_SELECT_CB
type PPlotSelect struct {
	Sender      *Handle
	Return      int32
	Index       int32
	SampleIndex int32
	X           float32
	Y           float32
	Select      int32
}

// Iup callback IUP_SELECTBEGIN_CB
type PPlotSelectBegin struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_SELECTEND_CB
type PPlotSelectEnd struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_EDIT_CB
type PPlotEdit struct {
	Sender      *Handle
	Return      int32
	Index       int32
	SampleIndex int32
	X           float32
	Y           float32
	NewX        *float32
	NewY        *float32
}

// Iup callback IUP_EDITBEGIN_CB
type PPlotEditBegin struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_EDITEND_CB
type PPlotEditEnd struct {
	Sender *Handle
	Return int32
}

// Iup callback IUP_PREDRAW_CB
type PPlotPreDraw struct {
	Sender *Handle
	Return int32
	Canvas unsafe.Pointer
}

// Iup callback IUP_POSTDRAW_CB
type PPlotPostDraw struct {
	Sender *Handle
	Return int32
	Canvas unsafe.Pointer
}

// Iup callback IUP_COMPLETED_CB
type WebBrowserCompleted struct {
	Sender *Handle
	Return int32
	Url    string
}

// Iup callback IUP_ERROR_CB
type WebBrowserError struct {
	Sender *Handle
	Return int32
	Url    string
}

// Iup callback IUP_NAVIGATE_CB
type WebBrowserNavigate struct {
	Sender *Handle
	Return int32
	Url    string
}

// Iup callback IUP_NEWWINDOW_CB
type WebBrowserNewWindow struct {
	Sender *Handle
	Return int32
	Url    string
}
