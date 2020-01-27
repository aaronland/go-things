window.addEventListener("load", function load(event){

    var pl = document.getElementById("place");

    var query = document.getElementById("query");    
    var results = document.getElementById("results");
    var thing_form = document.getElementById("thing-form");
    var thing_label = document.getElementById("thing-label");
    var thing_body = document.getElementById("thing-body");
    var thing_submit = document.getElementById("thing-submit");        

    thing_submit.onclick = function(){
    
	var depicts_id = thing_body.getAttribute("data-whosonfirst-id");
	var text = thing_body.value;

	console.log(depicts_id, text);

	try {
	    
	    var on_success = function(rsp){
		console.log("OKAY", rsp);
	    };
	    
	    var on_error = function(err){
		console.log("ERROR", err);
	    };
	    
	    things.api.add(depicts_id, text, on_success, on_error);
	    
	} catch (e){

	    console.log("SAD", e);
	}

	return false;
    };
    
    var draw_thing = function(name, id){

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

	for (var i=0; i < count; i++){

	    var r = results[i];
	    
	    var id = r.id;
	    var name = r.name;
	    var placetype = r.placetype;
	    
	    var label = name;

	    var lineage = r.lineage;
	    lineage = lineage[0];

	    try {
		var region = lineage.region;
		label = label + ", " + region.name;		
	    } catch (e) {
		// console.log("REGION", e);
	    }
	    
	    try {
		var country = lineage.country;
		label = label + ", " + country.name;
	    } catch (e) {
		// console.log("COUNTRY", e);
	    }

	    label = label + " (" + placetype + ")";
	    
	    var item = document.createElement("li");
	    item.setAttribute("data-whosonfirst-id", id);
	    item.appendChild(document.createTextNode(label));

	    item.onclick = function(){
		draw_thing(name, id);
		return false;
	    };
	    
	    list.appendChild(item);
	}

	var results = document.getElementById("results");
	results.innerHTML = "";

	results.appendChild(list);
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

	things.placeholder.client.search(text, on_success, on_error);
    };

});