$(function () {
    let containers = new Set();

    function addContainer() {
        let val = parseInt($("#containerInput").val(), 10);
        if (!val || val <= 1) {
            alert("Container size must be > 1");
            return;
        }
        if (containers.has(val)) {
            alert("Container already added");
            return;
        }
        containers.add(val);

        $("#containerList").append(
            '<span class="badge bg-primary me-1">' + val + '</span>' +
            '<input type="hidden" name="containers" value="' + val + '">'
        );
        $("#containerInput").val("");
    }

    $("#addContainer").click(addContainer);

    $("#containerInput").keydown(function(e) {
        if (e.key === "Enter" || e.keyCode === 13) {
            e.preventDefault(); // prevent form submission
            addContainer();
        }
    });

    $("#packerForm").submit(function (e) {
        let goods = parseInt($("#goodsInput").val(), 10);
        if (containers.size === 0) {
            alert("Add at least one container");
            e.preventDefault();
            return;
        }
        if (!goods || goods <= 0) {
            alert("Goods must be > 0");
            e.preventDefault();
            return;
        }
    });

    $("#runTests").click(function () {
        window.location.href = "/tests";
    });
});
