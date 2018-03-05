// ** variable "data" is pulled in from input.js **

formFields = ["id", "name", "title", "birth", "death", "bio", "notes"]
// search all the people to find where in the array
// a particular ID lives
function findIndexByID(id) {
    for(let i=0; i < data.people.length; i++) {
        if(data.people[i].id == id) return i;
    }
    return 0;
}

// pretty-print the full family's JSON
function printResult() {
    result = document.getElementById("result");
    result.innerHTML = JSON.stringify(data, null, 4);
}

// fetch a form field's current value
function getFormValue(id) {
    result = document.getElementById(id);
    return result.value;
}

// modify a form field's current value
function setFormValue(id, value) {
    if(value) { // don't write "undefined" in a bunch of places
        result = document.getElementById(id);
        result.value = value;
    }
}

// remove text from a form field
function clearFormValue(id) {
    result = document.getElementById(id);
    result.value = "";
}

// load the person with a specified ID into the form
function loadPersonToForm() {
    id = getFormValue("toLoad");
    for(var i=0; person = data.people[i]; i++) {
        if(person.id == id) {
            for (var field of formFields) {
                setFormValue(field, person[field]);
            }
        }
    }
    clearFormValue("toLoad");
}

function loadBackOneGeneration() {
    id = getFormValue("id");
    for(var i=0; person = data.people[i]; i++) {
        if(person.mother == id || person.father == id) {
            for (var field of formFields) {
                setFormValue(field, person[field]);
            }
        }
    }
}

// record the newly input person into the collection
function addPersonFromForm() {
    newPerson = {
        id: getFormValue("id"),
        name: getFormValue("name"),
        title: getFormValue("title"),
        birth: getFormValue("birth"),
        death: getFormValue("death"),
        bio: getFormValue("bio"),
        notes: getFormValue("notes")
    };
    data.people.push(newPerson);
    printResult();
}

function addRelation(relation) {
    max++;
    // find current person's position in the array:
    current = findIndexByID(getFormValue("id"));
    data.people[current][relation] = max;
    for (var field of formFields) {
        clearFormValue(field);
        console.log(field)
    }
    setFormValue("id", max);
    printResult();
}


// ****  STARTUP ****

// print out any pre-loaded people:
printResult()
// determine the highest ID in use by the people,
// so we have a starting point for new ones:
max = 0;
for(var i=0; i < data.people.length; i++) {
    if (data.people[i].id > max) {
        max = data.people[i].id;
    }
}
