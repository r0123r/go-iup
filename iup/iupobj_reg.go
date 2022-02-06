package iup

const (
	DIALOG        = "dialog"
	FILEDLG       = "filedlg"
	MESSAGEDLG    = "messagedlg"
	COLORDLG      = "colordlg"
	FONTDLG       = "fontdlg"
	FILL          = "fill"
	HBOX          = "hbox"
	VBOX          = "vbox"
	ZBOX          = "zbox"
	RADIO         = "radio"
	NORMALIZER    = "normalizer"
	CBOX          = "cbox"
	SBOX          = "sbox"
	SCROLLBOX     = "scrollbox"
	BACKGROUNDBOX = "backgroundbox"
	SPLIT         = "split"
	ITEM          = "item"
	MENU          = "menu"
	SEPARATOR     = "separator"
	SUBMENU       = "submenu"
	CLIPBOARD     = "clipboard"
	TIMER         = "timer"
	USER          = "user"
	BUTTON        = "button"
	CANVAS        = "canvas"
	FRAME         = "flatframe"
	LABEL         = "flatlabel"
	LIST          = "list"
	PROGRESSBAR   = "gauge" //"progressbar"
	SPIN          = "spin"
	SPINBOX       = "spinbox"
	TABS          = "flattabs"
	TEXT          = "text"
	TOGGLE        = "toggle"
	TREE          = "tree"
	VAL           = "val"
	CELLS         = "cells"
	COLORBAR      = "colorbar"
	COLORBROWSER  = "colorbrowser"
	DIAL          = "dial"
	MATRIX        = "matrix"
	OLECONTROL    = "olecontrol"
	PPLOT         = "plot"
	PROGRESSDLG   = "progressdlg"
	WEBBROWSER    = "webbrowser"
	TUIOCLIENT    = "tuioclient"
)

func RegisterAllClass() {
	RegisterClass("dialog", NewClassInfo("dialog", dialog_SetCallback))
	RegisterClass("filedlg", NewClassInfo("filedlg", filedlg_SetCallback))
	RegisterClass("messagedlg", NewClassInfo("messagedlg", messagedlg_SetCallback))
	RegisterClass("colordlg", NewClassInfo("colordlg", colordlg_SetCallback))
	RegisterClass("fontdlg", NewClassInfo("fontdlg", fontdlg_SetCallback))
	RegisterClass("progresslg", NewClassInfo("progressdlg", fontdlg_SetCallback))
	RegisterClass("fill", NewClassInfo("fill", fill_SetCallback))
	RegisterClass("hbox", NewClassInfo("hbox", hbox_SetCallback))
	RegisterClass("vbox", NewClassInfo("vbox", vbox_SetCallback))
	RegisterClass("zbox", NewClassInfo("zbox", zbox_SetCallback))
	RegisterClass("radio", NewClassInfo("radio", radio_SetCallback))
	RegisterClass("normalizer", NewClassInfo("normalizer", normalizer_SetCallback))
	RegisterClass("cbox", NewClassInfo("cbox", cbox_SetCallback))
	RegisterClass("sbox", NewClassInfo("sbox", sbox_SetCallback))
	RegisterClass("scrollbox", NewClassInfo("scrollbox", scrollbox_SetCallback))
	RegisterClass("backgroundbox", NewClassInfo("backgroundbox", backgroundbox_SetCallback))
	RegisterClass("split", NewClassInfo("split", split_SetCallback))
	RegisterClass("item", NewClassInfo("item", item_SetCallback))
	RegisterClass("menu", NewClassInfo("menu", menu_SetCallback))
	RegisterClass("separator", NewClassInfo("separator", separator_SetCallback))
	RegisterClass("submenu", NewClassInfo("submenu", submenu_SetCallback))
	RegisterClass("clipboard", NewClassInfo("clipboard", clipboard_SetCallback))
	RegisterClass("timer", NewClassInfo("timer", timer_SetCallback))
	RegisterClass("user", NewClassInfo("user", user_SetCallback))
	RegisterClass("button", NewClassInfo("button", button_SetCallback))
	RegisterClass("canvas", NewClassInfo("canvas", canvas_SetCallback))
	RegisterClass("frame", NewClassInfo("frame", frame_SetCallback))
	RegisterClass("label", NewClassInfo("label", label_SetCallback))
	RegisterClass("list", NewClassInfo("list", list_SetCallback))
	RegisterClass("progressbar", NewClassInfo("progressbar", progressbar_SetCallback))
	RegisterClass("spin", NewClassInfo("spin", spin_SetCallback))
	RegisterClass("spinbox", NewClassInfo("spinbox", spinbox_SetCallback))
	RegisterClass("tabs", NewClassInfo("tabs", tabs_SetCallback))
	RegisterClass("text", NewClassInfo("text", text_SetCallback))
	RegisterClass("toggle", NewClassInfo("toggle", toggle_SetCallback))
	RegisterClass("tree", NewClassInfo("tree", tree_SetCallback))
	RegisterClass("val", NewClassInfo("val", val_SetCallback))
	RegisterClass("cells", NewClassInfo("cells", cells_SetCallback))
	RegisterClass("colorbar", NewClassInfo("colorbar", colorbar_SetCallback))
	RegisterClass("colorbrowser", NewClassInfo("colorbrowser", colorbrowser_SetCallback))
	RegisterClass("dial", NewClassInfo("dial", dial_SetCallback))
	RegisterClass("matrix", NewClassInfo("matrix", matrix_SetCallback))
	RegisterClass("olecontrol", NewClassInfo("olecontrol", olecontrol_SetCallback))
	RegisterClass("plot", NewClassInfo("plot", pplot_SetCallback))
	RegisterClass("webbrowser", NewClassInfo("webbrowser", webbrowser_SetCallback))
	RegisterClass("tuioclient", NewClassInfo("tuioclient", tuioclient_SetCallback))
}

// Iup control Dialog
//
// Callback CLOSE_CB : func(arg *DialogClose)
//
// Callback RESIZE_CB : func(arg *DialogResize)
//
// Callback DROPFILES_CB : func(arg *DialogDropFiles)
//
// Callback SHOW_CB : func(arg *DialogShow)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Dialog(a ...interface{}) *Handle {
	return New(DIALOG, a...)
}

// Iup control FileDlg
//
// Callback FILE_CB : func(arg *FileDlgFile)
//
// Callback HELP_CB : func(arg *CommonHelp)
func FileDlg(a ...interface{}) *Handle {
	return New(FILEDLG, a...)
}

// Iup control MessageDlg
//
// Callback HELP_CB : func(arg *CommonHelp)
func MessageDlg(a ...interface{}) *Handle {
	return New(MESSAGEDLG, a...)
}

// Iup control ColorDlg
//
// Callback HELP_CB : func(arg *CommonHelp)
func ColorDlg(a ...interface{}) *Handle {
	return New(COLORDLG, a...)
}

// Iup control FontDlg
//
// Callback HELP_CB : func(arg *CommonHelp)
func FontDlg(a ...interface{}) *Handle {
	return New(FONTDLG, a...)
}

// Iup control ProgressDlg
//
// Callback HELP_CB : func(arg *CommonHelp)
func ProgressDlg(a ...interface{}) *Handle {
	return New(PROGRESSDLG, a...)
}

// Iup control Fill
//
//
func Fill(a ...interface{}) *Handle {
	return New(FILL, a...)
}

// Iup control Hbox
//
//
func Hbox(a ...interface{}) *Handle {
	return New(HBOX, a...)
}

// Iup control Vbox
//
//
func Vbox(a ...interface{}) *Handle {
	return New(VBOX, a...)
}

// Iup control Zbox
//
//
func Zbox(a ...interface{}) *Handle {
	return New(ZBOX, a...)
}

// Iup control Radio
//
//
func Radio(a ...interface{}) *Handle {
	return New(RADIO, a...)
}

// Iup control Normalizer
//
//
func Normalizer(a ...interface{}) *Handle {
	return New(NORMALIZER, a...)
}

// Iup control Cbox
//
//
func Cbox(a ...interface{}) *Handle {
	return New(CBOX, a...)
}

// Iup control Sbox
//
//
func Sbox(a ...interface{}) *Handle {
	return New(SBOX, a...)
}

// Iup control ScrollBox
//
//
func ScrollBox(a ...interface{}) *Handle {
	return New(SCROLLBOX, a...)
}

// Iup control BackgroundBox
//
//
func BackgroundBox(a ...interface{}) *Handle {
	return New(BACKGROUNDBOX, a...)
}

// Iup control Split
//
//
func Split(a ...interface{}) *Handle {
	return New(SPLIT, a...)
}

// Iup control Item
//
// Callback ACTION : func(arg *ItemAction)
//
// Callback HIGHLIGHT_CB : func(arg *ItemHighlight)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback HELP_CB : func(arg *CommonHelp)
func Item(a ...interface{}) *Handle {
	return New(ITEM, a...)
}

// Iup control Menu
//
// Callback OPEN_CB : func(arg *MenuOpen)
//
// Callback MENUCLOSE_CB : func(arg *MenuClose)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
func Menu(a ...interface{}) *Handle {
	return New(MENU, a...)
}

// Iup control Separator
//
//
func Separator(a ...interface{}) *Handle {
	return New(SEPARATOR, a...)
}

// Iup control SubMenu
//
// Callback HIGHLIGHT_CB : func(arg *SubMenuHighlight)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
func SubMenu(a ...interface{}) *Handle {
	return New(SUBMENU, a...)
}

// Iup control Clipboard
//
//
func Clipboard(a ...interface{}) *Handle {
	return New(CLIPBOARD, a...)
}

// Iup control Timer
//
// Callback ACTION : func(arg *TimerAction)
func Timer(a ...interface{}) *Handle {
	return New(TIMER, a...)
}

// Iup control User
//
//
func User(a ...interface{}) *Handle {
	return New(USER, a...)
}

// Iup control Button
//
// Callback ACTION : func(arg *ButtonAction)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Button(a ...interface{}) *Handle {
	return New(BUTTON, a...)
}

// Iup control Canvas
//
// Callback ACTION : func(arg *CanvasAction)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback DROPFILES_CB : func(arg *CanvasDropFiles)
//
// Callback FOCUS_CB : func(arg *CanvasFocus)
//
// Callback MOTION_CB : func(arg *MouseMotion)
//
// Callback KEYPRESS_CB : func(arg *CanvasKeyPress)
//
// Callback RESIZE_CB : func(arg *CanvasResize)
//
// Callback SCROLL_CB : func(arg *CanvasScroll)
//
// Callback TOUCH_CB : func(arg *TouchEvent)
//
// Callback MULTITOUCH_CB : func(arg *MultiTouchEvent)
//
// Callback WHEEL_CB : func(arg *CanvasWheel)
//
// Callback WOM_CB : func(arg *CanvasWom)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Canvas(a ...interface{}) *Handle {
	return New(CANVAS, a...)
}

// Iup control Frame
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
func Frame(a ...interface{}) *Handle {
	return New(FRAME, a...)
}

// Iup control Label
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback DROPFILES_CB : func(arg *LabelDropFiles)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
func Label(a ...interface{}) *Handle {
	return New(LABEL, a...)
}

// Iup control List
//
// Callback ACTION : func(arg *ListAction)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback CARET_CB : func(arg *ListCaret)
//
// Callback DBLCLICK_CB : func(arg *ListDblclick)
//
// Callback DROPDOWN_CB : func(arg *ListDropDown)
//
// Callback DROPFILES_CB : func(arg *ListDropFiles)
//
// Callback EDIT_CB : func(arg *ListEdit)
//
// Callback MOTION_CB : func(arg *MouseMotion)
//
// Callback MULTISELECT_CB : func(arg *ListMultiSelect)
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func List(a ...interface{}) *Handle {
	return New(LIST, a...)
}

// Iup control ProgressBar
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
func ProgressBar(a ...interface{}) *Handle {
	return New(PROGRESSBAR, a...)
}

// Iup control Spin
//
// Callback SPIN_CB : func(arg *SpinInc)
func Spin(a ...interface{}) *Handle {
	return New(SPIN, a...)
}

// Iup control SpinBox
//
// Callback SPIN_CB : func(arg *SpinBoxInc)
func SpinBox(a ...interface{}) *Handle {
	return New(SPINBOX, a...)
}

// Iup control Tabs
//
// Callback TABCHANGE_CB : func(arg *TabsChange)
//
// Callback TABCHANGEPOS_CB : func(arg *TabsChangePos)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Tabs(a ...interface{}) *Handle {
	return New(TABS, a...)
}

// Iup control Text
//
// Callback ACTION : func(arg *TextAction)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback CARET_CB : func(arg *TextCaret)
//
// Callback DROPFILES_CB : func(arg *TextDropFiles)
//
// Callback MOTION_CB : func(arg *MouseMotion)
//
// Callback SPIN_CB : func(arg *TextSpin)
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Text(a ...interface{}) *Handle {
	return New(TEXT, a...)
}

// Iup control Toggle
//
// Callback ACTION : func(arg *ToggleAction)
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Toggle(a ...interface{}) *Handle {
	return New(TOGGLE, a...)
}

// Iup control Tree
//
// Callback SELECTION_CB : func(arg *TreeSelection)
//
// Callback MULTISELECTION_CB : func(arg *TreeMultiSelection)
//
// Callback BRANCHOPEN_CB : func(arg *TreeBranchOpen)
//
// Callback BRANCHCLOSE_CB : func(arg *TreeBranchClose)
//
// Callback EXECUTELEAF_CB : func(arg *TreeExecuteLeaf)
//
// Callback SHOWRENAME_CB : func(arg *TreeShowRename)
//
// Callback RENAME_CB : func(arg *TreeRename)
//
// Callback DRAGDROP_CB : func(arg *TreeDragDrop)
//
// Callback NODEREMOVED_CB : func(arg *TreeNodeRemoved)
//
// Callback RIGHTCLICK_CB : func(arg *TreeRightClick)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback DROPFILES_CB : func(arg *TreeDropFiles)
//
// Callback MOTION_CB : func(arg *MouseMotion)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Tree(a ...interface{}) *Handle {
	return New(TREE, a...)
}

// Iup control Val
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Val(a ...interface{}) *Handle {
	return New(VAL, a...)
}

// Iup control Cells
//
// Callback DRAW_CB : func(arg *CellsDraw)
//
// Callback HEIGHT_CB : func(arg *CellsHeight)
//
// Callback HSPAN_CB : func(arg *CellsHspan)
//
// Callback MOUSECLICK_CB : func(arg *CellsMouseClick)
//
// Callback MOUSEMOTION_CB : func(arg *CellsMouseMotion)
//
// Callback NCOLS_CB : func(arg *CellsNcols)
//
// Callback NLINES_CB : func(arg *CellsNlines)
//
// Callback SCROLLING_CB : func(arg *CellsScrolling)
//
// Callback VSPAN_CB : func(arg *CellsVspan)
//
// Callback WIDTH_CB : func(arg *CellsWidth)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Cells(a ...interface{}) *Handle {
	return New(CELLS, a...)
}

// Iup control Colorbar
//
// Callback CELL_CB : func(arg *ColorbarCell)
//
// Callback EXTENDED_CB : func(arg *ColorbarExtended)
//
// Callback SELECT_CB : func(arg *ColorbarSelect)
//
// Callback SWITCH_CB : func(arg *ColorbarSwitch)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Colorbar(a ...interface{}) *Handle {
	return New(COLORBAR, a...)
}

// Iup control ColorBrowser
//
// Callback CHANGE_CB : func(arg *ColorBrowserChange)
//
// Callback DRAG_CB : func(arg *ColorBrowserDrag)
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func ColorBrowser(a ...interface{}) *Handle {
	return New(COLORBROWSER, a...)
}

// Iup control Dial
//
// Callback BUTTON_PRESS_CB : func(arg *DialButtonPress)
//
// Callback BUTTON_RELEASE_CB : func(arg *DialButtonRelease)
//
// Callback MOUSEMOVE_CB : func(arg *DialMouseMove)
//
// Callback VALUECHANGED_CB : func(arg *ValueChanged)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func Dial(a ...interface{}) *Handle {
	return New(DIAL, a...)
}

// Iup control Matrix
//
// Callback ACTION_CB : func(arg *MatrixAction)
//
// Callback CLICK_CB : func(arg *MatrixClick)
//
// Callback RELEASE_CB : func(arg *MatrixRelease)
//
// Callback MOUSEMOVE_CB : func(arg *MatrixMouseMove)
//
// Callback ENTERITEM_CB : func(arg *MatrixEnterItem)
//
// Callback LEAVEITEM_CB : func(arg *MatrixLeaveItem)
//
// Callback SCROLLTOP_CB : func(arg *MatrixScrollTop)
//
// Callback BGCOLOR_CB : func(arg *MatrixBgcolor)
//
// Callback FGCOLOR_CB : func(arg *MatrixFgcolor)
//
// Callback FONT_CB : func(arg *MatrixFont)
//
// Callback DRAW_CB : func(arg *MatrixDraw)
//
// Callback DROPCHECK_CB : func(arg *MatrixDropCheck)
//
// Callback DROP_CB : func(arg *MatrixDrop)
//
// Callback DROPSELECT_CB : func(arg *MatrixDropSelect)
//
// Callback EDITION_CB : func(arg *MatrixEdition)
//
// Callback VALUE_CB : func(arg *MatrixValue)
//
// Callback VALUE_EDIT_CB : func(arg *MatrixValueEdit)
//
// Callback MARK_CB : func(arg *MatrixMark)
//
// Callback MARKEDIT_CB : func(arg *MatrixMarkEdit)
//
// Callback ACTION : func(arg *CanvasAction)
//
// Callback SCROLL_CB : func(arg *CanvasScroll)
//
// Callback KEYPRESS_CB : func(arg *CanvasKeyPress)
//
// Callback MOTION_CB : func(arg *MouseMotion)
//
// Callback RESIZE_CB : func(arg *CanvasResize)
//
// Callback BUTTON_CB : func(arg *MouseButton)
//
// Callback MAP_CB : func(arg *CommonMap)
func Matrix(a ...interface{}) *Handle {
	return New(MATRIX, a...)
}

// Iup control OleControl
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
func OleControl(a ...interface{}) *Handle {
	return New(OLECONTROL, a...)
}

// Iup control PPlot
//
// Callback DELETE_CB : func(arg *PPlotDelete)
//
// Callback DELETEBEGIN_CB : func(arg *PPlotDeleteBegin)
//
// Callback DELETEEND_CB : func(arg *PPlotDeleteEnd)
//
// Callback SELECT_CB : func(arg *PPlotSelect)
//
// Callback SELECTBEGIN_CB : func(arg *PPlotSelectBegin)
//
// Callback SELECTEND_CB : func(arg *PPlotSelectEnd)
//
// Callback EDIT_CB : func(arg *PPlotEdit)
//
// Callback EDITBEGIN_CB : func(arg *PPlotEditBegin)
//
// Callback EDITEND_CB : func(arg *PPlotEditEnd)
//
// Callback PREDRAW_CB : func(arg *PPlotPreDraw)
//
// Callback POSTDRAW_CB : func(arg *PPlotPostDraw)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func PPlot(a ...interface{}) *Handle {
	return New(PPLOT, a...)
}

// Iup control WebBrowser
//
// Callback COMPLETED_CB : func(arg *WebBrowserCompleted)
//
// Callback ERROR_CB : func(arg *WebBrowserError)
//
// Callback NAVIGATE_CB : func(arg *WebBrowserNavigate)
//
// Callback NEWWINDOW_CB : func(arg *WebBrowserNewWindow)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func WebBrowser(a ...interface{}) *Handle {
	return New(WEBBROWSER, a...)
}

// Iup control TuioClient
//
// Callback TOUCH_CB : func(arg *TouchEvent)
//
// Callback MULTITOUCH_CB : func(arg *MultiTouchEvent)
//
// Callback DEFAULT_ACTION : func(arg *CommonDefaultAction)
//
// Callback MAP_CB : func(arg *CommonMap)
//
// Callback UNMAP_CB : func(arg *CommonUnmap)
//
// Callback DESTROY_CB : func(arg *CommonDestroy)
//
// Callback GETFOCUS_CB : func(arg *CommonGetFocus)
//
// Callback KILLFOCUS_CB : func(arg *CommonKillFocus)
//
// Callback ENTERWINDOW_CB : func(arg *CommonEnterWindow)
//
// Callback LEAVEWINDOW_CB : func(arg *CommonLeaveWindow)
//
// Callback HELP_CB : func(arg *CommonHelp)
//
// Callback K_ANY : func(arg *CommonKeyAny)
func TuioClient(a ...interface{}) *Handle {
	return New(TUIOCLIENT, a...)
}
