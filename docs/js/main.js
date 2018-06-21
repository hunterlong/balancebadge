

$("#create_badge").click(function() {

    var address = $("#address").val();
    var coin = $("#coin").find("option:selected").val();
    var svg = "https://img.balancebadge.io/"+coin+"/"+address+".svg";

    var url = "https://etherscan.io/address/"+address;

    var html = "<a href=\""+url+"\"><img src=\""+svg+"\"></a>";
    var markup = "[![Balance]("+svg+")]("+url+")";
    var restruct = ".. image:: "+svg+" :alt: Balance";
    var ascii = "image:"+svg+"[Balance]";

    $("#html").text(html);
    $("#markup").text(markup);
    $("#restruct").text(restruct);
    $("#ascii").text(ascii);

    $("#render_img").removeClass('d-none');

    return false;
});