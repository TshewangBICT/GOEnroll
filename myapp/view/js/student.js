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
    const students = JSON.parse(data)
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

function getFormData() {
    var formData = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    return formData
}

function addStudent() {
    var data = getFormData()

    fetch("/student/add", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => {
        // check the response from fetch is resolved or rejected
        if (response1.ok) {
            var sid = data.stdid;
            // student/1001
            fetch("/student/"+sid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
        } else {
            throw new Error(response1.status)
        }
    }).catch(e => {
        if (e.message == 401) {
            alert("User Not Logged in")
            window.open("index.html", "_self")
        } else if (e.message == 400) {
            alert("Bad Request")
            window.open("index.html", "_self")
        } else {
            alert("Internal Server Error")
        }
    });
    
    resetform();
} 

var selectedRow = null;
function updateStudent(input) {
    // get the selected row
    selectedRow = input.parentElement.parentElement

    document.getElementById("sid").value = selectedRow.cells[0].innerHTML
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML
    document.getElementById("email").value = selectedRow.cells[3].innerHTML

    sid = selectedRow.cells[0].innerHTML
    // change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "Update"
    btn.setAttribute("onclick", "updateAPIRequest(sid)")
}

function updateAPIRequest(oldSid) {
    // send call update API
    var newData = getFormData()
    fetch("/student/"+oldSid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(res => {
        if (res.ok) {
            // fill in the selected row with updated values
            selectedRow.cells[0].innerHTML = newData.stdid
            selectedRow.cells[1].innerHTML = newData.fname
            selectedRow.cells[2].innerHTML = newData.lname
            selectedRow.cells[3].innerHTML = newData.email

            // change the button vlaue to initial state
            var btn = document.getElementById("button-add")
            btn.innerHTML = "update"
            btn.setAttribute("onclick", "addStudent()")

            selectedRow = null;

            resetform();

        } else {
            alert("Server: Update request error")
        }
    });
}

function deleteStudent(r) {
    if (confirm("Are you sure you want to DELETE this student?")) {
        selectedRow = r.parentElement.parentElement;
        sid = selectedRow.cells[0].innerHTML;

        fetch("student/"+sid, {
            method: "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
        });
        var rowIndex = selectedRow.rowIndex; //index starts from 0
        if (rowIndex > 0) {
            document.getElementById("myTable").deleteRow(rowIndex);
        }
        selectedRow = null;
    }
}