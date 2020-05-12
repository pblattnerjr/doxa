//$(document).ready(function(){
window.onload = function () {	//code here...
	var code = $(".codemirror-textarea")[0];
	var editor = CodeMirror.fromTextArea(code, {
		mode: "lml",
		theme: "twilight",
		lineNumbers: true,
		lineWrapping: true,
		extraKeys: {
			  "Ctrl-1": showCss
			, "Ctrl-2": showTks
			, "Ctrl-3": showLmlPaths
			, "Ctrl-4": showKeywords},
		foldGutter: true,
		viewportMargin: Infinity,
		gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"]
	});
	editor.foldCode(CodeMirror.Pos(0, 0));
	editor.setSize("100%", "100%");

	var snippets = []; // will be set to keywords, lmlPaths, tks depending on which ctrl key pressed

	var keywords = [
		{ text: 'insert ', displayText: 'insert: add contents of specified template.' },
		{ text: 'nid ', displayText: 'nid: no id. Insert text as is.' },
		{ text: 'rid ', displayText: 'rid: insert text using relative ID.' },
		{ text: 'sid ', displayText: 'sid: insert text using specific ID.' },
		{ text: 'ver ', displayText: 'ver: display version name.' },
		{ text: 'when_exists rid  use:  {\n\totherwise use:\n}', displayText: 'when_exists' },
	];
	function showKeywords(cm) {
		snippets = keywords;
		snippet(cm);
	}
	function showCss(cm) {
		snippets = css;
		snippet(cm);
	}
	function showTks(cm) {
		snippets = tks;
		snippet(cm);
	}
	function showLmlPaths(cm) {
		snippets = lmlPaths;
		snippet(cm);
	}
	function snippet(cm) {
		// this gets called everytime a key is pressed (after the shortcut was called)
		CodeMirror.showHint(cm, function () {
			var cursor = cm.getCursor();
			var token = cm.getTokenAt(cursor);
			var currentWord = token.string;
			var start = token.start;
			var end = cursor.ch;
			var line = cursor.line;
			var list = snippets.filter(function (item) {
				return item.text.indexOf(currentWord) >= 0;
			});
			return {
				list: list.length ? list : snippets,
				from: CodeMirror.Pos(line, start),
				to: CodeMirror.Pos(line, end)
			};
		}, { completeSingle: false });
	}
	var lmlPaths = [
		{ text: '"ages/blocks/LI/_01_Enarxis__/daily" ', displayText: 'ages/blocks/LI/_01_Enarxis__/daily' },
		{ text: '"ages/blocks/LI/_01_Enarxis__/vesperal" ', displayText: 'ages/blocks/LI/_01_Enarxis__/vesperal' },
		{ text: '"ages/blocks/LN/GreatLitany01" ', displayText: 'ages/blocks/LN/GreatLitany01' },
		{ text: '"ages/blocks/LN/GreatLitany02" ', displayText: 'ages/blocks/LN/GreatLitany02' },
		{ text: '"ages/blocks/LN/ShortLitany00" ', displayText: 'ages/blocks/LN/ShortLitany00' },
	];
	var tks = [ // topic-keys
		{ text: '"actors~Deacon" ', displayText: 'actors~Deacon' },
		{ text: '"actors~People" ', displayText: 'actors~People' },
		{ text: '"actors~Priest" ', displayText: 'actors~Priest' },
		{ text: '"eu.lichrysbasil~euLI.Key0109.text" ', displayText: 'eu.lichrysbasil~euLI.Key0109.text' },
		{ text: '"eu.lichrysbasil~euLI.Key0200.text" ', displayText: 'eu.lichrysbasil~euLI.Key0200.text' },
		{ text: '"eu.lichrysbasil~euLI.Key0201.text" ', displayText: 'eu.lichrysbasil~euLI.Key0201.text' },
		{ text: '"oc.*~ocMO.Ode41.text" ', displayText: 'oc.*~ocMO.Ode41.text' },
		{ text: '"oc.*~ocMO.Ode42.text" ', displayText: 'oc.*~ocMO.Ode42.text' },
		{ text: '"oc.*~ocMO.Ode42.text" ', displayText: 'oc.*~ocMO.Ode42.text' },
		{ text: '"oc.*~ocMO.Ode42.text" ', displayText: 'oc.*~ocMO.Ode42.text' },
		{ text: '"oc.*~ocMO.Ode43.text" ', displayText: 'oc.*~ocMO.Ode43.text' },
		{ text: '"prayers~Doxa" ', displayText: 'prayers~Doxa' },
		{ text: '"prayers~enarxis02" ', displayText: 'prayers~enarxis02' },
		{ text: '"prayers~exc01" ', displayText: 'prayers~exc01' },
		{ text: '"prayers~pet02" ', displayText: 'prayers~pet02' },
		{ text: '"prayers~res04p" ', displayText: 'prayers~res04p' },
		{ text: '"prayers~res04p" ', displayText: 'prayers~res04p' },
		{ text: '"prayers~verse_Trinity.text" ', displayText: 'prayers~verse_Trinity.text' },
	];
	// you created the json using cssscanner/main.go in this project.
	var css = JSON.parse('[{"text":".IndexLink ","DisplayText":"IndexLink"},{"text":".actor ","DisplayText":"actor"},{"text":".agesMenu ","DisplayText":"agesMenu"},{"text":".ages_menu_link ","DisplayText":"ages_menu_link"},{"text":".alttext ","DisplayText":"alttext"},{"text":".alwb_media_dropdown_div ","DisplayText":"alwb_media_dropdown_div"},{"text":".anum ","DisplayText":"anum"},{"text":".bigblack ","DisplayText":"bigblack"},{"text":".black ","DisplayText":"black"},{"text":".bmc_collapse ","DisplayText":"bmc_collapse"},{"text":".bold ","DisplayText":"bold"},{"text":".bolditalics ","DisplayText":"bolditalics"},{"text":".bolditalicsred ","DisplayText":"bolditalicsred"},{"text":".boldred ","DisplayText":"boldred"},{"text":".books_index_table ","DisplayText":"books_index_table"},{"text":".break ","DisplayText":"break"},{"text":".byzscore ","DisplayText":"byzscore"},{"text":".chant ","DisplayText":"chant"},{"text":".chapverse ","DisplayText":"chapverse"},{"text":".clockbox ","DisplayText":"clockbox"},{"text":".container ","DisplayText":"container"},{"text":".content ","DisplayText":"content"},{"text":".cover1 ","DisplayText":"cover1"},{"text":".cover2 ","DisplayText":"cover2"},{"text":".cover3 ","DisplayText":"cover3"},{"text":".cover4 ","DisplayText":"cover4"},{"text":".cover5 ","DisplayText":"cover5"},{"text":".credits1 ","DisplayText":"credits1"},{"text":".credits2 ","DisplayText":"credits2"},{"text":".credits3 ","DisplayText":"credits3"},{"text":".custom_index_table ","DisplayText":"custom_index_table"},{"text":".designation ","DisplayText":"designation"},{"text":".dialog ","DisplayText":"dialog"},{"text":".dialogzero ","DisplayText":"dialogzero"},{"text":".em ","DisplayText":"em"},{"text":".emc_collapse ","DisplayText":"emc_collapse"},{"text":".fa ","DisplayText":"fa"},{"text":".fa_arrows ","DisplayText":"fa_arrows"},{"text":".fa_bars ","DisplayText":"fa_bars"},{"text":".fa_byznote ","DisplayText":"fa_byznote"},{"text":".fa_calendar ","DisplayText":"fa_calendar"},{"text":".fa_caret_square_o_left ","DisplayText":"fa_caret_square_o_left"},{"text":".fa_caret_square_o_right ","DisplayText":"fa_caret_square_o_right"},{"text":".fa_columns ","DisplayText":"fa_columns"},{"text":".fa_envelope ","DisplayText":"fa_envelope"},{"text":".fa_info_circle ","DisplayText":"fa_info_circle"},{"text":".fa_list_alt ","DisplayText":"fa_list_alt"},{"text":".fa_minus ","DisplayText":"fa_minus"},{"text":".fa_mobile ","DisplayText":"fa_mobile"},{"text":".fa_money ","DisplayText":"fa_money"},{"text":".fa_music ","DisplayText":"fa_music"},{"text":".fa_plus ","DisplayText":"fa_plus"},{"text":".fa_question_circle ","DisplayText":"fa_question_circle"},{"text":".fa_text_height ","DisplayText":"fa_text_height"},{"text":".fa_undo ","DisplayText":"fa_undo"},{"text":".fa_volume_up ","DisplayText":"fa_volume_up"},{"text":".glyphicon ","DisplayText":"glyphicon"},{"text":".heirmos ","DisplayText":"heirmos"},{"text":".horizontalLine ","DisplayText":"horizontalLine"},{"text":".hymn ","DisplayText":"hymn"},{"text":".hymnlinefirst ","DisplayText":"hymnlinefirst"},{"text":".hymnlinelast ","DisplayText":"hymnlinelast"},{"text":".hymnlinemiddle ","DisplayText":"hymnlinemiddle"},{"text":".inaudible ","DisplayText":"inaudible"},{"text":".index_books_file_link ","DisplayText":"index_books_file_link"},{"text":".index_books_file_link_td ","DisplayText":"index_books_file_link_td"},{"text":".index_books_file_timestamp_td ","DisplayText":"index_books_file_timestamp_td"},{"text":".index_books_file_type_td ","DisplayText":"index_books_file_type_td"},{"text":".index_books_language ","DisplayText":"index_books_language"},{"text":".index_books_language_td ","DisplayText":"index_books_language_td"},{"text":".index_commemoration ","DisplayText":"index_commemoration"},{"text":".index_content ","DisplayText":"index_content"},{"text":".index_custom_file_link ","DisplayText":"index_custom_file_link"},{"text":".index_custom_file_link_td ","DisplayText":"index_custom_file_link_td"},{"text":".index_custom_file_timestamp_td ","DisplayText":"index_custom_file_timestamp_td"},{"text":".index_custom_file_type_td ","DisplayText":"index_custom_file_type_td"},{"text":".index_custom_language ","DisplayText":"index_custom_language"},{"text":".index_custom_language_td ","DisplayText":"index_custom_language_td"},{"text":".index_day_link ","DisplayText":"index_day_link"},{"text":".index_day_td ","DisplayText":"index_day_td"},{"text":".index_day_tr ","DisplayText":"index_day_tr"},{"text":".index_desc_1 ","DisplayText":"index_desc_1"},{"text":".index_desc_2 ","DisplayText":"index_desc_2"},{"text":".index_desc_3 ","DisplayText":"index_desc_3"},{"text":".index_file_link ","DisplayText":"index_file_link"},{"text":".index_file_link_td ","DisplayText":"index_file_link_td"},{"text":".index_file_timestamp ","DisplayText":"index_file_timestamp"},{"text":".index_file_timestamp_td ","DisplayText":"index_file_timestamp_td"},{"text":".index_month ","DisplayText":"index_month"},{"text":".index_month_td ","DisplayText":"index_month_td"},{"text":".index_month_tr ","DisplayText":"index_month_tr"},{"text":".index_service_day ","DisplayText":"index_service_day"},{"text":".index_service_day_td ","DisplayText":"index_service_day_td"},{"text":".index_service_day_tr ","DisplayText":"index_service_day_tr"},{"text":".index_service_language ","DisplayText":"index_service_language"},{"text":".index_service_language_td ","DisplayText":"index_service_language_td"},{"text":".index_service_language_tr ","DisplayText":"index_service_language_tr"},{"text":".index_title ","DisplayText":"index_title"},{"text":".index_title_date ","DisplayText":"index_title_date"},{"text":".italics ","DisplayText":"italics"},{"text":".italicsblack ","DisplayText":"italicsblack"},{"text":".italicsred ","DisplayText":"italicsred"},{"text":".leftCell ","DisplayText":"leftCell"},{"text":".m_g_e ","DisplayText":"m_g_e"},{"text":".mediaMenuItemId ","DisplayText":"mediaMenuItemId"},{"text":".mediaMenuItemPeople ","DisplayText":"mediaMenuItemPeople"},{"text":".media_group ","DisplayText":"media_group"},{"text":".media_group_empty ","DisplayText":"media_group_empty"},{"text":".media_icon ","DisplayText":"media_icon"},{"text":".melody ","DisplayText":"melody"},{"text":".missalreading ","DisplayText":"missalreading"},{"text":".missaltitle ","DisplayText":"missaltitle"},{"text":".mixed ","DisplayText":"mixed"},{"text":".mode ","DisplayText":"mode"},{"text":".name ","DisplayText":"name"},{"text":".noprintactor ","DisplayText":"noprintactor"},{"text":".noprintdesig ","DisplayText":"noprintdesig"},{"text":".noprintprayer ","DisplayText":"noprintprayer"},{"text":".noprintrub ","DisplayText":"noprintrub"},{"text":".notavailable ","DisplayText":"notavailable"},{"text":".numright ","DisplayText":"numright"},{"text":".optional ","DisplayText":"optional"},{"text":".panelTitle ","DisplayText":"panelTitle"},{"text":".panel_body ","DisplayText":"panel_body"},{"text":".panel_heading ","DisplayText":"panel_heading"},{"text":".percentage ","DisplayText":"percentage"},{"text":".pixels ","DisplayText":"pixels"},{"text":".plain ","DisplayText":"plain"},{"text":".point ","DisplayText":"point"},{"text":".prayer ","DisplayText":"prayer"},{"text":".prayerzero ","DisplayText":"prayerzero"},{"text":".pref_closer ","DisplayText":"pref_closer"},{"text":".pref_instructions ","DisplayText":"pref_instructions"},{"text":".pref_left ","DisplayText":"pref_left"},{"text":".pref_panel ","DisplayText":"pref_panel"},{"text":".pref_right ","DisplayText":"pref_right"},{"text":".pref_row_spacer ","DisplayText":"pref_row_spacer"},{"text":".pref_spacer ","DisplayText":"pref_spacer"},{"text":".print_btn ","DisplayText":"print_btn"},{"text":".reading ","DisplayText":"reading"},{"text":".readingcenter ","DisplayText":"readingcenter"},{"text":".readingcenterzero ","DisplayText":"readingcenterzero"},{"text":".readingzero ","DisplayText":"readingzero"},{"text":".red ","DisplayText":"red"},{"text":".reference ","DisplayText":"reference"},{"text":".rubric ","DisplayText":"rubric"},{"text":".small ","DisplayText":"small"},{"text":".smallcenter ","DisplayText":"smallcenter"},{"text":".source ","DisplayText":"source"},{"text":".source0 ","DisplayText":"source0"},{"text":".source1 ","DisplayText":"source1"},{"text":".strong ","DisplayText":"strong"},{"text":".titlezero ","DisplayText":"titlezero"},{"text":".typikon ","DisplayText":"typikon"},{"text":".typikonblock ","DisplayText":"typikonblock"},{"text":".typikoncenter ","DisplayText":"typikoncenter"},{"text":".typikoncenteromitted ","DisplayText":"typikoncenteromitted"},{"text":".typikoncommemoration ","DisplayText":"typikoncommemoration"},{"text":".typikonindent ","DisplayText":"typikonindent"},{"text":".typikonnote ","DisplayText":"typikonnote"},{"text":".typikonomitted ","DisplayText":"typikonomitted"},{"text":".undefined ","DisplayText":"undefined"},{"text":".verse ","DisplayText":"verse"},{"text":".versecenter ","DisplayText":"versecenter"},{"text":".versiondesignation ","DisplayText":"versiondesignation"},{"text":".verticalLine ","DisplayText":"verticalLine"}]');
};
