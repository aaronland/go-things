var things = things || {};
things.placeholder = things.placeholder || {};

things.placeholder.client = (function(){

    var endpoint = 'http://localhost:8181/api';

    var self = {
	
	'init': function(){

	},
	
	'search': function(text, on_success, on_error){

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
		    var data = JSON.parse(raw);

		    data = self._filter_results(data);
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
	    
	    var uri = endpoint + "/parser/search/?term=" + text;
	    console.log("FETCH", uri);
	    
	    try {

		console.log("GET",uri);
		var req = new XMLHttpRequest();
		
		req.addEventListener("load", on_load);
		req.addEventListener("error", on_failed);
		req.addEventListener("abort", on_abort);
		
		req.open("GET", uri, true);
		req.send();
		
	    } catch (e) {
		on_error(e);
		return false;
	    }
	    
	},

	'_filter_results': function(results){

	    var filtered = [];

	    var count = results.length;

	    for (var i=0;i < count; i++){

		var r = results[i];
		
		switch (r.placetype) {
		    case "locality":
		    filtered.push(r);
		    break;
		    case "neighbourhood":
		    filtered.push(r);
		    default:
		    continue;
		}
	    }

	    return filtered;
	}
	
    };

    return self;
    
})();
