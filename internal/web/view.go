package web

var mainJS = `
{{define "mainJS"}}
	var ws = new WebSocket("ws://" + window.location.host + "/memstats-feed")
	var tpl = _.template(document.getElementById("ms-viewer-template").innerHTML)

	// SOCKET /memstats-feeds
	ws.onopen = function () {

		// ON MESSAGE /memstats-feed
		ws.onmessage = function (evt) {
			var memdata = JSON.parse(evt.data);
			var humanized = _.clone(memdata);
			
			console.log(memdata);
			[ // Convert byte values to readable form.
				"Alloc", "TotalAlloc", "Sys", "HeapAlloc", "HeapSys", "HeapIdle",
				"HeapInuse", "HeapReleased", "StackInuse", "StackSys", "MSpanInuse",
				"MSpanSys", "MCacheInuse", "MCacheSys", "NextGC"

			].forEach(function (key) {
				humanized.Stats[key] = bytesToSize(memdata.Stats[key]); 
			}); 

			document.getElementById("ms-viewer").innerHTML = tpl(humanized);
		}

		// ON CLOSE /memstats-feeds
		ws.onclose = function () {
			console.log("MEMSTAT: Disconnected.")
		}
	}

	// Converts bytes to human-readable form with precision(3)
	function bytesToSize(bytes) {
		if(bytes == 0) return '0 Byte';
		var k = 1000, i = Math.floor(Math.log(bytes) / Math.log(k));
		var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
		return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
	};
{{end}}`

var stylesheet = `
{{define "stylesheet"}}
	div.group {
		width: 250px;
		padding: 20px;
		float: left;
		border: 1px solid #dfdfdf;
	}

	div.group div.cell {
		margin: 5px 0 0 0;
	}

	div.group h4 {
		margin: 5px 0 0 0;
	}
{{end}}`
