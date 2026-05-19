window.onload = function() {
    loadStudents();
    loadCourses();
    loadEnrollments();
}

function loadStudents() {
    fetch("/student/all")
    .then(response => response.json())
    .then(students => {
        const select = document.getElementById("sid");
        select.innerHTML = '<option value="">Select Student</option>';
        students.forEach(student => {
            select.innerHTML += `<option value="${student.stdid}">${student.stdid}</option>`;
        });
    })
    .catch(error => console.error("Error loading students:", error));
}

function loadCourses() {
    fetch("/course/all")
    .then(response => response.json())
    .then(courses => {
        const select = document.getElementById("cid");
        select.innerHTML = '<option value="">Select Course</option>';
        courses.forEach(course => {
            select.innerHTML += `<option value="${course.cid}">${course.cid}</option>`;
        });
    })
    .catch(error => console.error("Error loading courses:", error));
}

function addEnroll() {
    const stdid = document.getElementById("sid").value;
    const cid = document.getElementById("cid").value;
    
    if (!stdid || !cid) {
        alert("Please select both student and course");
        return;
    }
    
    fetch("/enroll", {
        method: "POST",
        body: JSON.stringify({ stdid: parseInt(stdid), cid: cid }),
        headers: { "Content-type": "application/json" }
    })
    .then(response => {
        if (response.status === 201) {
            alert("Enrolled successfully!");
            document.getElementById("sid").value = "";
            document.getElementById("cid").value = "";
            loadEnrollments();
        } else if (response.status === 403) {
            alert("Student already enrolled in this course!");
        } else {
            alert("Enrollment failed");
        }
    })
    .catch(error => console.error("Error:", error));
}

function loadEnrollments() {
    fetch("/enroll/all")
    .then(response => response.json())
    .then(enrollments => {
        const table = document.getElementById("myTable");
        while(table.rows.length > 1) table.deleteRow(1);
        
        enrollments.forEach(enroll => {
            const row = table.insertRow();
            row.insertCell(0).innerHTML = enroll.StdId;
            row.insertCell(1).innerHTML = enroll.cid;
            row.insertCell(2).innerHTML = enroll.date;
            row.insertCell(3).innerHTML = `<input type="button" onClick="deleteEnroll(this)" value="delete" id="button-1" />`;
        });
    })
    .catch(error => console.error("Error loading enrollments:", error));
}

function deleteEnroll(button) {
    if (confirm("Are you sure you want to delete this enrollment?")) {
        const row = button.parentElement.parentElement;
        const stdid = row.cells[0].innerHTML;
        const cid = row.cells[1].innerHTML;
        
        fetch(`/enroll/${stdid}/${cid}`, { method: "DELETE" })
        .then(response => {
            if (response.ok) {
                row.remove();
                alert("Enrollment deleted successfully!");
            } else {
                alert("Delete failed");
            }
        })
        .catch(error => console.error("Error:", error));
    }
}