
#include <stdlib.h>
#include <string.h>
#include <iup.h>
#define GOIUP "_GOIUP_"

// IUP definitions not defined
#define IUP_UNMAP_CB        "UNMAP_CB"
#define IUP_DESTROY_CB      "DESTROY_CB"
#define IUP_CARET_CB        "CARET_CB"
#define IUP_DBLCLICK_CB     "DBLCLICK_CB"
#define IUP_EDIT_CB         "EDIT_CB"
#define IUP_MULTISELECT_CB  "MULTISELECT_CB"
#define IUP_VALUECHANGED_CB "VALUECHANGED_CB"
#define IUP_TABCHANGE_CB    "TABCHANGE_CB"
#define IUP_TABCHANGEPOS_CB "TABCHANGEPOS_CB"
#define IUP_SPIN_CB         "SPIN_CB"
#define IUP_FILE_CB			"FILE_CB"
#define IUP_FOCUS_CB		"FOCUS_CB"
#define IUP_TOUCH_CB		"TOUCH_CB"
#define IUP_MULTITOUCH_CB	"MULTITOUCH_CB"
#define IUP_DROPDOWN_CB		"DROPDOWN_CB"
#define IUP_MULTISELECTION_CB "MULTISELECTION_CB"
#define IUP_SHOWRENAME_CB	"SHOWRENAME_CB"
#define IUP_RENAME_CB		"RENAME_CB"
#define IUP_DRAGDROP_CB		"DRAGDROP_CB"
#define IUP_NODEREMOVED_CB	"NODEREMOVED_CB"
#define IUP_BUTTON_PRESS_CB	"BUTTON_PRESS_CB"
#define IUP_BUTTON_RELEASE_CB "BUTTON_RELEASE_CB"
#define IUP_ACTION_CB		"ACTION_CB"
#define IUP_RELEASE_CB		"RELEASE_CB"
#define IUP_SCROLLTOP_CB	"SCROLLTOP_CB"
#define IUP_BGCOLOR_CB		"BGCOLOR_CB"
#define IUP_FGCOLOR_CB		"FGCOLOR_CB"
#define IUP_FONT_CB			"FONT_CB"
#define IUP_MARK_CB			"MARK_CB"
#define IUP_MARKEDIT_CB		"MARKEDIT_CB"
#define IUP_DELETE_CB		"DELETE_CB"
#define IUP_DELETEBEGIN_CB	"DELETEBEGIN_CB"
#define IUP_DELETEEND_CB	"DELETEEND_CB"
#define IUP_SELECTBEGIN_CB	"SELECTBEGIN_CB"
#define IUP_SELECTEND_CB	"SELECTEND_CB"
#define IUP_EDITBEGIN_CB	"EDITBEGIN_CB"
#define	IUP_EDITEND_CB		"EDITEND_CB"
#define IUP_PREDRAW_CB		"PREDRAW_CB"
#define IUP_POSTDRAW_CB		"POSTDRAW_CB"
#define IUP_COMPLETED_CB	"COMPLETED_CB"
#define IUP_ERROR_CB		"ERROR_CB"
#define IUP_NAVIGATE_CB		"NAVIGATE_CB"
#define IUP_NEWWINDOW_CB	"NEWWINDOW_CB"
#define IUP_PLOTMOTION_CB	"PLOTMOTION_CB"
#define IUP_XTICKFORMATNUMBER_CB "XTICKFORMATNUMBER_CB"
#define IUP_EXECUTEBRANCH_CB "EXECUTEBRANCH_CB"
// extern int goCommonDefaultAction(void*);
// static void iupSetCommonDefaultAction(Ihandle* ih)
// {
// 	IupSetCallback(ih,IUP_DEFAULT_ACTION,(Icallback)&goCommonDefaultAction);
// }

extern int goCommonMap(void*);
static void iupSetCommonMap(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MAP_CB,(Icallback)&goCommonMap);
}

extern int goCommonUnmap(void*);
static void iupSetCommonUnmap(Ihandle* ih)
{
	IupSetCallback(ih,IUP_UNMAP_CB,(Icallback)&goCommonUnmap);
}

extern int goCommonDestroy(void*);
static void iupSetCommonDestroy(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DESTROY_CB,(Icallback)&goCommonDestroy);
}

extern int goCommonGetFocus(void*);
static void iupSetCommonGetFocus(Ihandle* ih)
{
	IupSetCallback(ih,IUP_GETFOCUS_CB,(Icallback)&goCommonGetFocus);
}

extern int goCommonKillFocus(void*);
static void iupSetCommonKillFocus(Ihandle* ih)
{
	IupSetCallback(ih,IUP_KILLFOCUS_CB,(Icallback)&goCommonKillFocus);
}

extern int goCommonEnterWindow(void*);
static void iupSetCommonEnterWindow(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ENTERWINDOW_CB,(Icallback)&goCommonEnterWindow);
}

extern int goCommonLeaveWindow(void*);
static void iupSetCommonLeaveWindow(Ihandle* ih)
{
	IupSetCallback(ih,IUP_LEAVEWINDOW_CB,(Icallback)&goCommonLeaveWindow);
}

extern int goCommonHelp(void*);
static void iupSetCommonHelp(Ihandle* ih)
{
	IupSetCallback(ih,IUP_HELP_CB,(Icallback)&goCommonHelp);
}

extern int goCommonKeyAny(void*,int);
static void iupSetCommonKeyAny(Ihandle* ih)
{
	IupSetCallback(ih,IUP_K_ANY,(Icallback)&goCommonKeyAny);
}

extern int goDialogClose(void*);
static void iupSetDialogClose(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CLOSE_CB,(Icallback)&goDialogClose);
}

extern int goDialogResize(void*,int,int);
static void iupSetDialogResize(Ihandle* ih)
{
	IupSetCallback(ih,IUP_RESIZE_CB,(Icallback)&goDialogResize);
}

extern int goDialogDropFiles(void*,char*,int,int,int);
static void iupSetDialogDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goDialogDropFiles);
}

extern int goDialogShow(void*,int);
static void iupSetDialogShow(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SHOW_CB,(Icallback)&goDialogShow);
}

extern int goFileDlgFile(void*,char*,char*);
static void iupSetFileDlgFile(Ihandle* ih)
{
	IupSetCallback(ih,IUP_FILE_CB,(Icallback)&goFileDlgFile);
}

extern int goItemAction(void*);
static void iupSetItemAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goItemAction);
}

extern int goItemHighlight(void*);
static void iupSetItemHighlight(Ihandle* ih)
{
	IupSetCallback(ih,IUP_HIGHLIGHT_CB,(Icallback)&goItemHighlight);
}

extern int goMenuOpen(void*);
static void iupSetMenuOpen(Ihandle* ih)
{
	IupSetCallback(ih,IUP_OPEN_CB,(Icallback)&goMenuOpen);
}

extern int goMenuClose(void*);
static void iupSetMenuClose(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MENUCLOSE_CB,(Icallback)&goMenuClose);
}

extern int goSubMenuHighlight(void*);
static void iupSetSubMenuHighlight(Ihandle* ih)
{
	IupSetCallback(ih,IUP_HIGHLIGHT_CB,(Icallback)&goSubMenuHighlight);
}

extern int goTimerAction(void*);
static void iupSetTimerAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION_CB,(Icallback)&goTimerAction);
}

extern int goButtonAction(void*);
static void iupSetButtonAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goButtonAction);
}

extern int goMouseButton(void*,int,int,int,int,char*);
static void iupSetMouseButton(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BUTTON_CB,(Icallback)&goMouseButton);
}

extern int goCanvasAction(void*);
static void iupSetCanvasAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goCanvasAction);
}

extern int goCanvasDropFiles(void*,char*,int,int,int);
static void iupSetCanvasDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goCanvasDropFiles);
}

extern int goCanvasFocus(void*,int);
static void iupSetCanvasFocus(Ihandle* ih)
{
	IupSetCallback(ih,IUP_FOCUS_CB,(Icallback)&goCanvasFocus);
}

extern int goMouseMotion(void*,int,int,char*);
static void iupSetMouseMotion(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MOTION_CB,(Icallback)&goMouseMotion);
}

extern int goPlotMotion(void*,double,double,char*);
static void iupSetPlotMotion(Ihandle* ih)
{
	IupSetCallback(ih,IUP_PLOTMOTION_CB,(Icallback)&goPlotMotion);
}

extern int goPlotXtick(void*,char*,char*,double,char*);
static void iupSetPlotXtick(Ihandle* ih)
{
	IupSetCallback(ih,IUP_XTICKFORMATNUMBER_CB,(Icallback)&goPlotXtick);
}

extern int goCanvasKeyPress(void*,int,int);
static void iupSetCanvasKeyPress(Ihandle* ih)
{
	IupSetCallback(ih,IUP_KEYPRESS_CB,(Icallback)&goCanvasKeyPress);
}

extern int goCanvasResize(void*,int,int);
static void iupSetCanvasResize(Ihandle* ih)
{
	IupSetCallback(ih,IUP_RESIZE_CB,(Icallback)&goCanvasResize);
}

extern int goCanvasScroll(void*,int,float,float);
static void iupSetCanvasScroll(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SCROLL_CB,(Icallback)&goCanvasScroll);
}

extern int goTouchEvent(void*,int,int,char*);
static void iupSetTouchEvent(Ihandle* ih)
{
	IupSetCallback(ih,IUP_TOUCH_CB,(Icallback)&goTouchEvent);
}

extern int goMultiTouchEvent(void*,int,int*,int*,int*,int*);
static void iupSetMultiTouchEvent(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MULTITOUCH_CB,(Icallback)&goMultiTouchEvent);
}

extern int goCanvasWheel(void*,float,int,int,char*);
static void iupSetCanvasWheel(Ihandle* ih)
{
	IupSetCallback(ih,IUP_WHEEL_CB,(Icallback)&goCanvasWheel);
}

extern int goCanvasWom(void*,int);
static void iupSetCanvasWom(Ihandle* ih)
{
	IupSetCallback(ih,IUP_WOM_CB,(Icallback)&goCanvasWom);
}

extern int goLabelDropFiles(void*,char*,int,int,int);
static void iupSetLabelDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goLabelDropFiles);
}

extern int goListAction(void*,char*,int,int);
static void iupSetListAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goListAction);
}

extern int goListCaret(void*,int,int,int);
static void iupSetListCaret(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CARET_CB,(Icallback)&goListCaret);
}

extern int goListDblclick(void*,int,char*);
static void iupSetListDblclick(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DBLCLICK_CB,(Icallback)&goListDblclick);
}

extern int goListDropDown(void*,int);
static void iupSetListDropDown(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPDOWN_CB,(Icallback)&goListDropDown);
}

extern int goListDropFiles(void*,char*,int,int,int);
static void iupSetListDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goListDropFiles);
}

extern int goListEdit(void*,int,char*);
static void iupSetListEdit(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EDIT_CB,(Icallback)&goListEdit);
}

extern int goListMultiSelect(void*,char*);
static void iupSetListMultiSelect(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MULTISELECT_CB,(Icallback)&goListMultiSelect);
}

extern int goValueChanged(void*);
static void iupSetValueChanged(Ihandle* ih)
{
	IupSetCallback(ih,IUP_VALUECHANGED_CB,(Icallback)&goValueChanged);
}

extern int goSpinInc(void*,int);
static void iupSetSpinInc(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SPIN_CB,(Icallback)&goSpinInc);
}

extern int goSpinBoxInc(void*,int);
static void iupSetSpinBoxInc(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SPIN_CB,(Icallback)&goSpinBoxInc);
}

extern int goTabsChange(void*,void*,void*);
static void iupSetTabsChange(Ihandle* ih)
{
	IupSetCallback(ih,IUP_TABCHANGE_CB,(Icallback)&goTabsChange);
}

extern int goTabsChangePos(void*,int,int);
static void iupSetTabsChangePos(Ihandle* ih)
{
	IupSetCallback(ih,IUP_TABCHANGEPOS_CB,(Icallback)&goTabsChangePos);
}

extern int goTextAction(void*,int,char*);
static void iupSetTextAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goTextAction);
}

extern int goTextCaret(void*,int,int,int);
static void iupSetTextCaret(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CARET_CB,(Icallback)&goTextCaret);
}

extern int goTextDropFiles(void*,char*,int,int,int);
static void iupSetTextDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goTextDropFiles);
}

extern int goTextSpin(void*,int);
static void iupSetTextSpin(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SPIN_CB,(Icallback)&goTextSpin);
}

extern int goToggleAction(void*,int);
static void iupSetToggleAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION,(Icallback)&goToggleAction);
}

extern int goTreeSelection(void*,int,int);
static void iupSetTreeSelection(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SELECTION_CB,(Icallback)&goTreeSelection);
}

extern int goTreeMultiSelection(void*,int*,int);
static void iupSetTreeMultiSelection(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MULTISELECTION_CB,(Icallback)&goTreeMultiSelection);
}

extern int goTreeBranchOpen(void*,int);
static void iupSetTreeBranchOpen(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BRANCHOPEN_CB,(Icallback)&goTreeBranchOpen);
}

extern int goTreeBranchClose(void*,int);
static void iupSetTreeBranchClose(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BRANCHCLOSE_CB,(Icallback)&goTreeBranchClose);
}

extern int goTreeExecuteLeaf(void*,int);
static void iupSetTreeExecuteLeaf(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EXECUTELEAF_CB,(Icallback)&goTreeExecuteLeaf);
}

extern int goTreeExecuteBranch(void*,int);
static void iupSetTreeExecuteBranch(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EXECUTEBRANCH_CB,(Icallback)&goTreeExecuteBranch);
}

extern int goTreeShowRename(void*,int);
static void iupSetTreeShowRename(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SHOWRENAME_CB,(Icallback)&goTreeShowRename);
}

extern int goTreeRename(void*,int,char*);
static void iupSetTreeRename(Ihandle* ih)
{
	IupSetCallback(ih,IUP_RENAME_CB,(Icallback)&goTreeRename);
}

extern int goTreeDragDrop(void*,int,int,int,int);
static void iupSetTreeDragDrop(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DRAGDROP_CB,(Icallback)&goTreeDragDrop);
}

extern int goTreeNodeRemoved(void*,void*);
static void iupSetTreeNodeRemoved(Ihandle* ih)
{
	IupSetCallback(ih,IUP_NODEREMOVED_CB,(Icallback)&goTreeNodeRemoved);
}

extern int goTreeRightClick(void*,int);
static void iupSetTreeRightClick(Ihandle* ih)
{
	IupSetCallback(ih,IUP_RIGHTCLICK_CB,(Icallback)&goTreeRightClick);
}

extern int goTreeDropFiles(void*,char*,int,int,int);
static void iupSetTreeDropFiles(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPFILES_CB,(Icallback)&goTreeDropFiles);
}

extern int goCellsDraw(void*,int,int,int,int,int,int,void*);
static void iupSetCellsDraw(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DRAW_CB,(Icallback)&goCellsDraw);
}

extern int goCellsHeight(void*,int);
static void iupSetCellsHeight(Ihandle* ih)
{
	IupSetCallback(ih,IUP_HEIGHT_CB,(Icallback)&goCellsHeight);
}

extern int goCellsHspan(void*,int,int);
static void iupSetCellsHspan(Ihandle* ih)
{
	IupSetCallback(ih,IUP_HSPAN_CB,(Icallback)&goCellsHspan);
}

extern int goCellsMouseClick(void*,int,int,int,int,int,int,char*);
static void iupSetCellsMouseClick(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MOUSECLICK_CB,(Icallback)&goCellsMouseClick);
}

extern int goCellsMouseMotion(void*,int,int,int,int,char*);
static void iupSetCellsMouseMotion(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MOUSEMOTION_CB,(Icallback)&goCellsMouseMotion);
}

extern int goCellsNcols(void*);
static void iupSetCellsNcols(Ihandle* ih)
{
	IupSetCallback(ih,IUP_NCOLS_CB,(Icallback)&goCellsNcols);
}

extern int goCellsNlines(void*);
static void iupSetCellsNlines(Ihandle* ih)
{
	IupSetCallback(ih,IUP_NLINES_CB,(Icallback)&goCellsNlines);
}

extern int goCellsScrolling(void*,int,int);
static void iupSetCellsScrolling(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SCROLLING_CB,(Icallback)&goCellsScrolling);
}

extern int goCellsVspan(void*,int,int);
static void iupSetCellsVspan(Ihandle* ih)
{
	IupSetCallback(ih,IUP_VSPAN_CB,(Icallback)&goCellsVspan);
}

extern int goCellsWidth(void*,int);
static void iupSetCellsWidth(Ihandle* ih)
{
	IupSetCallback(ih,IUP_WIDTH_CB,(Icallback)&goCellsWidth);
}

extern char* goColorbarCell(void*,int);
static void iupSetColorbarCell(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CELL_CB,(Icallback)&goColorbarCell);
}

extern int goColorbarExtended(void*,int);
static void iupSetColorbarExtended(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EXTENDED_CB,(Icallback)&goColorbarExtended);
}

extern int goColorbarSelect(void*,int,int);
static void iupSetColorbarSelect(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SELECT_CB,(Icallback)&goColorbarSelect);
}

extern int goColorbarSwitch(void*,int,int);
static void iupSetColorbarSwitch(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SWITCH_CB,(Icallback)&goColorbarSwitch);
}

extern int goColorBrowserChange(void*,unsigned char,unsigned char,unsigned char);
static void iupSetColorBrowserChange(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CHANGE_CB,(Icallback)&goColorBrowserChange);
}

extern int goColorBrowserDrag(void*,unsigned char,unsigned char,unsigned char);
static void iupSetColorBrowserDrag(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DRAG_CB,(Icallback)&goColorBrowserDrag);
}

extern int goDialButtonPress(void*,double);
static void iupSetDialButtonPress(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BUTTON_PRESS_CB,(Icallback)&goDialButtonPress);
}

extern int goDialButtonRelease(void*,double);
static void iupSetDialButtonRelease(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BUTTON_RELEASE_CB,(Icallback)&goDialButtonRelease);
}

extern int goDialMouseMove(void*,double);
static void iupSetDialMouseMove(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MOUSEMOVE_CB,(Icallback)&goDialMouseMove);
}

extern int goMatrixAction(void*,int,int,int,int,char*);
static void iupSetMatrixAction(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ACTION_CB,(Icallback)&goMatrixAction);
}

extern int goMatrixClick(void*,int,int,char*);
static void iupSetMatrixClick(Ihandle* ih)
{
	IupSetCallback(ih,IUP_CLICK_CB,(Icallback)&goMatrixClick);
}

extern int goMatrixRelease(void*,int,int,char*);
static void iupSetMatrixRelease(Ihandle* ih)
{
	IupSetCallback(ih,IUP_RELEASE_CB,(Icallback)&goMatrixRelease);
}

extern int goMatrixMouseMove(void*,int,int);
static void iupSetMatrixMouseMove(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MOUSEMOVE_CB,(Icallback)&goMatrixMouseMove);
}

extern int goMatrixEnterItem(void*,int,int);
static void iupSetMatrixEnterItem(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ENTERITEM_CB,(Icallback)&goMatrixEnterItem);
}

extern int goMatrixLeaveItem(void*,int,int);
static void iupSetMatrixLeaveItem(Ihandle* ih)
{
	IupSetCallback(ih,IUP_LEAVEITEM_CB,(Icallback)&goMatrixLeaveItem);
}

extern int goMatrixScrollTop(void*,int,int);
static void iupSetMatrixScrollTop(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SCROLLTOP_CB,(Icallback)&goMatrixScrollTop);
}

extern int goMatrixBgcolor(void*,int,int,unsigned int*,unsigned int*,unsigned int*);
static void iupSetMatrixBgcolor(Ihandle* ih)
{
	IupSetCallback(ih,IUP_BGCOLOR_CB,(Icallback)&goMatrixBgcolor);
}

extern int goMatrixFgcolor(void*,int,int,unsigned int*,unsigned int*,unsigned int*);
static void iupSetMatrixFgcolor(Ihandle* ih)
{
	IupSetCallback(ih,IUP_FGCOLOR_CB,(Icallback)&goMatrixFgcolor);
}

extern int goMatrixFont(void*,int,int);
static void iupSetMatrixFont(Ihandle* ih)
{
	IupSetCallback(ih,IUP_FONT_CB,(Icallback)&goMatrixFont);
}

extern int goMatrixDraw(void*,int,int,int,int,int,int,void*);
static void iupSetMatrixDraw(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DRAW_CB,(Icallback)&goMatrixDraw);
}

extern int goMatrixDropCheck(void*,int,int);
static void iupSetMatrixDropCheck(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPCHECK_CB,(Icallback)&goMatrixDropCheck);
}

extern int goMatrixDrop(void*,void*,int,int);
static void iupSetMatrixDrop(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROP_CB,(Icallback)&goMatrixDrop);
}

extern int goMatrixDropSelect(void*,int,int,void*,char*,int,int);
static void iupSetMatrixDropSelect(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DROPSELECT_CB,(Icallback)&goMatrixDropSelect);
}

extern int goMatrixEdition(void*,int,int,int,int);
static void iupSetMatrixEdition(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EDITION_CB,(Icallback)&goMatrixEdition);
}

extern int goMatrixValue(void*,int,int);
static void iupSetMatrixValue(Ihandle* ih)
{
	IupSetCallback(ih,IUP_VALUE_CB,(Icallback)&goMatrixValue);
}

extern int goMatrixValueEdit(void*,int,int,char*);
static void iupSetMatrixValueEdit(Ihandle* ih)
{
	IupSetCallback(ih,IUP_VALUE_EDIT_CB,(Icallback)&goMatrixValueEdit);
}

extern int goMatrixMark(void*,int,int);
static void iupSetMatrixMark(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MARK_CB,(Icallback)&goMatrixMark);
}

extern int goMatrixMarkEdit(void*,int,int,int);
static void iupSetMatrixMarkEdit(Ihandle* ih)
{
	IupSetCallback(ih,IUP_MARKEDIT_CB,(Icallback)&goMatrixMarkEdit);
}

extern int goPPlotDelete(void*,int,int,float,float);
static void iupSetPPlotDelete(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DELETE_CB,(Icallback)&goPPlotDelete);
}

extern int goPPlotDeleteBegin(void*);
static void iupSetPPlotDeleteBegin(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DELETEBEGIN_CB,(Icallback)&goPPlotDeleteBegin);
}

extern int goPPlotDeleteEnd(void*);
static void iupSetPPlotDeleteEnd(Ihandle* ih)
{
	IupSetCallback(ih,IUP_DELETEEND_CB,(Icallback)&goPPlotDeleteEnd);
}

extern int goPPlotSelect(void*,int,int,float,float,int);
static void iupSetPPlotSelect(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SELECT_CB,(Icallback)&goPPlotSelect);
}

extern int goPPlotSelectBegin(void*);
static void iupSetPPlotSelectBegin(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SELECTBEGIN_CB,(Icallback)&goPPlotSelectBegin);
}

extern int goPPlotSelectEnd(void*);
static void iupSetPPlotSelectEnd(Ihandle* ih)
{
	IupSetCallback(ih,IUP_SELECTEND_CB,(Icallback)&goPPlotSelectEnd);
}

extern int goPPlotEdit(void*,int,int,float,float,float*,float*);
static void iupSetPPlotEdit(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EDIT_CB,(Icallback)&goPPlotEdit);
}

extern int goPPlotEditBegin(void*);
static void iupSetPPlotEditBegin(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EDITBEGIN_CB,(Icallback)&goPPlotEditBegin);
}

extern int goPPlotEditEnd(void*);
static void iupSetPPlotEditEnd(Ihandle* ih)
{
	IupSetCallback(ih,IUP_EDITEND_CB,(Icallback)&goPPlotEditEnd);
}

extern int goPPlotPreDraw(void*,void*);
static void iupSetPPlotPreDraw(Ihandle* ih)
{
	IupSetCallback(ih,IUP_PREDRAW_CB,(Icallback)&goPPlotPreDraw);
}

extern int goPPlotPostDraw(void*,void*);
static void iupSetPPlotPostDraw(Ihandle* ih)
{
	IupSetCallback(ih,IUP_POSTDRAW_CB,(Icallback)&goPPlotPostDraw);
}

extern int goWebBrowserCompleted(void*,char*);
static void iupSetWebBrowserCompleted(Ihandle* ih)
{
	IupSetCallback(ih,IUP_COMPLETED_CB,(Icallback)&goWebBrowserCompleted);
}

extern int goWebBrowserError(void*,char*);
static void iupSetWebBrowserError(Ihandle* ih)
{
	IupSetCallback(ih,IUP_ERROR_CB,(Icallback)&goWebBrowserError);
}

extern int goWebBrowserNavigate(void*,char*);
static void iupSetWebBrowserNavigate(Ihandle* ih)
{
	IupSetCallback(ih,IUP_NAVIGATE_CB,(Icallback)&goWebBrowserNavigate);
}

extern int goWebBrowserNewWindow(void*,char*);
static void iupSetWebBrowserNewWindow(Ihandle* ih)
{
	IupSetCallback(ih,IUP_NEWWINDOW_CB,(Icallback)&goWebBrowserNewWindow);
}


