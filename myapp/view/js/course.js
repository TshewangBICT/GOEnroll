window.onload = function() {
    fetch("course/all")
    .then(response => response.text())
    .then(data => showCourses(data))
    .catch(error => console.error("Error loading courses:", error))
}

function newRow(table, course) {
    var row = table.insertRow(table.length)

    var td = []
    for(i=0; i<table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i)
    }

        // insert data in the td cells
        td[0].innerHTML = course.cid
        td[1].innerHTML = course.coursename
        td[2].innerHTML = `<input type="button" onClick="deleteCourse(this)" value="delete" id="button-1" />`
        td[3].innerHTML = `<input type="button" onClick="updateCourse(this)" value="edit" id="button-2" />`
}

function showCourses(data) {
    const courses = JSON.parse(data)
    var table = document.getElementById("myTable");
    
    // Clear existing rows (except header)
    while(table.rows.length > 1) {
        table.deleteRow(1);
    }
    
    courses.forEach(course => {
        newRow(table, course)
    });  
}

function showCourse(data) {
    // console.log(data);
    // convert json string to js obj
    const course = JSON.parse(data)
    var table = document.getElementById("myTable")
    var row = table.insertRow(table.length)
    newRow(table, course)
}

// helper function to reset a form fields
function resetform() {
    document.getElementById("cid").value = "";
    document.getElementById("cname").value = "";
}

function getFormData() {
    var formData = {
        cid : document.getElementById("cid").value,
        coursename : document.getElementById("cname").value
    }
    return formData
}

function addCourse() {
    // create a javascript to store form data
    var data = getFormData()

    // Form validation
    if (data.cid === "") {
        alert("Course ID cannot be empty")
        return
    }
    if (data.coursename === "") {
        alert("Course name cannot be empty")
        return
    }

    //call POST API
    fetch("/course/add", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        // check the response from fetch is resolved or rejected
        if (response.ok) {
            var cid = data.cid
            return fetch("/course/"+cid)
        } else {
            throw new Error(response.statusText)
        }
    }).then(response => response.text())
    .then(data => showCourse(data))
    .catch(e => alert(e));

    resetform();
} 

var selectedRow = null;

function updateCourse(input) {
    // get the selected row
    selectedRow = input.parentElement.parentElement

    document.getElementById("cid").value = selectedRow.cells[0].innerHTML
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML

    var cid = selectedRow.cells[0].innerHTML
    // change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onclick", "updateAPIRequest('"+cid+"')")
}

function updateAPIRequest(oldCid) {
    // send call update API
    var newData = getFormData()
    
    fetch("/course/" + oldCid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok) {
            // fill in the selected row with updated values
            selectedRow.cells[0].innerHTML = newData.cid
            selectedRow.cells[1].innerHTML = newData.coursename

            // change the button value to initial state
            var btn = document.getElementById("button-add")
            btn.innerHTML = "Add"
            btn.setAttribute("onclick", "addCourse()")

            selectedRow = null;
            resetform();
        } else {
            alert("Server: Update request error")
        }
    });
}


function deleteCourse(input) {
    if (confirm("Are you sure you want to delete this course?")) {
        var selectedRow = input.parentElement.parentElement
        var cid = selectedRow.cells[0].innerHTML
        
        fetch("/course/" + cid, {
            method: "DELETE"
        }).then(response => {
            if (response.ok) {
                // Remove the row from the table
                selectedRow.remove()
            } else {
                throw new Error(response.statusText)
            }
        }).catch(e => alert(e))
    }
}

