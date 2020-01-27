var things = things || {};

things.api = (function(){

    var endpoint = location.origin;
    
    var self = {

	'add': function(depicts_id, text, on_success, on_error){

	    var form_data = new FormData();
	    form_data["id"] = depicts_id;
	    form_data["text"]= text;
	    
	    var on_load = function(rsp){

		var target = rsp.target;
		
		if (target.readyState != 4){
		    return;
		}
		
		var status_code = target['status'];
		var status_text = target['statusText'];
		
		if ((status_code < 200) || (status_code > 299)){
		    on_error({'code': status_code, 'message': status_text});
		    return;
		}
		
		try {
		    var raw = target['responseText'];
		    console.log("RAW", raw);
		    
		    var data = JSON.parse(raw);
		    console.log("DATA", data);
		    
		    on_success(data);
		} catch (e) {
		    on_error(e);
		}
		
	    };
	    
	    var on_failed = function(e){
		on_error(e);
	    };
	    
	    var on_abort = function(){
		on_error(e);
	    };
	    
	    var uri = endpoint + "/add/";

	    try {

		console.log("GET",uri);
		var req = new XMLHttpRequest();
		
		req.addEventListener("load", on_load);
		req.addEventListener("error", on_failed);
		req.addEventListener("abort", on_abort);
		
		req.open("POST", uri, true);
		req.send(form_data);
		
	    } catch (e) {
		on_error(e);
		return false;
	    }
	    
	},
    };

    return self;
    
})();
