window.addEventListener("load", function load(event){

    var pl = document.getElementById("place");

    var query = document.getElementById("query");    
    var results = document.getElementById("results");
    var thing_form = document.getElementById("thing-form");
    var thing_label = document.getElementById("thing-label");
    var thing_body = document.getElementById("thing-body");
    var thing_submit = document.getElementById("thing-submit");        
    var thing_cancel = document.getElementById("thing-cancel");

    var reset_form = function(){

	thing_body.removeAttribute("data-whosonfirst-id");
	thing_body.value = "";
	thing_label.innerText = "";

	thing_form.style.display = "none";
	
	results.innerHTML = "";
	results.style.display = "none";
	
	query.style.display = "block";
	pl.value = "";	
    };
    
    thing_cancel.onclick = function(){
	reset_form();
	return false;
    }
    
    thing_submit.onclick = function(){
    
	var depicts_id = thing_body.getAttribute("data-whosonfirst-id");
	depicts_id = parseInt(depicts_id);
	
	var text = thing_body.value;

	try {
	    
	    var on_success = function(rsp){
		console.log("OKAY", rsp);		
		reset_form();
	    };
	    
	    var on_error = function(err){
		console.log("ERROR", err);		
		reset_form();		
	    };

	    things.api.add(depicts_id, text, on_success, on_error);
	    
	} catch (e){
	    console.log("SAD", e);
	    reset_form();
	}

	return false;
    };
    
    var draw_thing = function(name, id){

	query.style.display = "none";
	results.style.display = "none";
	thing_form.style.display = "block";

	thing_label.innerText = "Things I like about " + name;
	
	thing_body.setAttribute("data-whosonfirst-id", id);
	// thing_body.innerHTML = "";
	
	console.log("THING", id);
	
    };
    
    var draw_results = function(results){

	var count = results.length;
	
	var list = document.createElement("ul");
	list.setAttribute("class", "select-menu");
	
	for (var i=0; i < count; i++){

	    var r = results[i];
	    
	    var id = r.id;
	    var name = r.name;
	    var placetype = r.placetype;
	    
	    var label = [];

	    var lineage = r.lineage;
	    lineage = lineage[0];

	    if (lineage.neighbourhood){
		label.push(" (" + lineage.neighbourhood.name + ")");
	    }
	    	    
	    if (lineage.macrohood){
		label.push(" (" + lineage.macrohood.name + ")");
	    }
	    
	    if (lineage.locality){
		label.push(lineage.locality.name);
	    }

	    if (lineage.county){
		label.push(" (" + lineage.county.name + ")");
	    }
	    
	    if (lineage.region){
		label.push(lineage.region.name);		
	    }

	    if (lineage.country){
		label.push(lineage.country.name);
	    }

	    var str_label = label.join(", ");

	    var item = document.createElement("li");
	    item.setAttribute("data-whosonfirst-id", id);
	    item.setAttribute("data-whosonfirst-name", str_label);    
	    item.appendChild(document.createTextNode(str_label));

	    item.onclick = function(e){

		var el = e.target;
		var id = el.getAttribute("data-whosonfirst-id");
		var name = el.getAttribute("data-whosonfirst-name");
		draw_thing(name, id);
		return false;
	    };
	    
	    list.appendChild(item);
	}

	var results = document.getElementById("results");
	results.innerHTML = "";

	results.appendChild(list);
	results.style.display = "block";
    };
    
    pl.oninput = function(e) {

	var el = e.target;
	var text = el.value;

	if (text.length < 3){
	    return;
	}
	
	var on_success = function(rsp){
	    draw_results(rsp);
	};

	var on_error = function(err){
	    console.log("SAD", err);
	};

	try {
	    things.placeholder.client.search(text, on_success, on_error);
	} catch (e){
	    console.log("WTF",e);
	}
    };

});
