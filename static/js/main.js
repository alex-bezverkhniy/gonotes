//(function(){
    var DONE = 4; // readyState 4 means the request is done.
    var OK = 200; // status 200 is a successful return.

    var scope = {
        note: {}
    };

    console.log("Init data");
    httpRequest = new XMLHttpRequest();

    httpRequest.open('GET', '/notes');
    httpRequest.send(null);

    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === DONE) {
            if (httpRequest.status === OK) {
                console.log(httpRequest.responseText);
                scope.notes = JSON.parse(httpRequest.responseText);
                render();
            }
        }        
    };

    // Data binding
    var elements = document.querySelectorAll("[data-tw-bind]");

    elements.forEach(function(element){
        if(element.type === 'text' || element.type === 'textarea') {
            var propToBind = element.getAttribute('data-tw-bind');
            element.onkeyup = function() {
                scope.note[propToBind] = element.value;
            };
        }
    });

    // Show json
    function render() {
        var el = document.querySelector("#json");
        el.innerHTML = httpRequest.responseText;

        var container = document.querySelector("#container");
        scope.notes.forEach(function(note) {
            var template = `<ul>
                <li>${note.title}</li>
            </ul>`;
            container.insertAdjacentHTML('beforeend', template);  
        });
        
    }

    // Get Note by ID

    function findByID(id) {
        httpRequest.open('GET', '/notes/'+id);
        httpRequest.send(null);
    }

    // Add note
    function addNote() {        
        httpRequest.open('POST', '/notes');
        httpRequest.send(JSON.stringify(scope.note));
    }

//})();