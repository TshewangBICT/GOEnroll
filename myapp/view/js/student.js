window.onload = function() {
    fetch("/student/all")
    .then(response => response.text())
    .then(data => showStudents(data))
}

function newRow(table, student) {
    var row = table.insertRow(table.length)

    var td = []
    for(i=0; i<table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i)
    }

        // insert data in the td cells
        td[0].innerHTML = student.stdid
        td[1].innerHTML = student.fname
        td[2].innerHTML = student.lname
        td[3].innerHTML = student.email
        td[4].innerHTML = `<input type="button" onClick="deleteStudent(this)" value="delete" id="button-1" />`
        td[5].innerHTML = `<input type="button" onClick="updateStudent(this)" value="edit" id="button-2" />`
}

function showStudents(data) {
    var students = JSON.parse(data)
    var table = document.getElementById("myTable");
    students.forEach(stud => {
        newRow(table, stud)
    });
    
}

function showStudent(data) {
    // console.log(data);
    // convert json string to js obj
    const student = JSON.parse(data)
    var table = document.getElementById("myTable")
    var row = table.insertRow(table.length)
    newRow(table, student)
    
}

// helper function to reset a form fields
function resetform() {
    document.getElementById("sid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("email").value = "";
}

function addStudent() {
    // create a javascript to store form data
    var data = {
        stdid: parseInt(document.getElementById("sid").value),
        fname: document.getElementById("fname").value,
        lname: document.getElementById("lname").value,
        email: document.getElementById("email").value
    }
    // form validation
    var sid = data.stdid
    if (isNaN(sid)) {
        alert("Enter a valid student ID")
        return
    } else if (data.email == "") {
        alert("Email cannot be empty")
        return
    } else if (data.fname == "")  {
        alert("Email cannot be empty")
        return
    }

    //call POST API
    //axios, fetch - to make http request
    //API route, req obj
    fetch("/student/add", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => {
        // check the response from fetch is resolved or rejected
        if (response1.ok) {
            // student/1001
            fetch("/student/"+sid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
        } else {
            throw new Error(response1.statusText)
        }
    }).catch(e => alert(e));
    resetform();
} 

function updateStudent(input) {
    // get the selected row
    var selectedRow = input.parentElement.parentElement
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML
    document.getElementById("email").value = selectedRow.cells[3].innerHTML

    var sid = selectedRow.cells[0].innerHTML
    // change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onClick", "updateAPIRequest(sid)")
}


