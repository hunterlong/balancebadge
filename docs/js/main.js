

$("#create_badge").click(function() {
    var address = $("#address").val();
    var coin = $("#coin").find("option:selected").val();
    if (address=="" || coin=="") {
        return false;
    }
    var svg = "https://img.balancebadge.io/"+coin+"/"+address+".svg";
    var usd_svg = "https://img.balancebadge.io/"+coin+"/"+address+"/usd.svg";
    var color_svg = "https://img.balancebadge.io/"+coin+"/"+address+".svg?color=cyan";
    var html = "<a href=\""+svg+"\"><img src=\""+svg+"\"></a>";
    var markup = "[![Balance]("+svg+")]("+svg+")";
    var restruct = ".. image:: "+svg+" :alt: Balance";
    var ascii = "image:"+svg+"[Balance]";

    var page_html = "<a target=\"_blank\" href=\""+svg+"\"><img src=\""+svg+"\"></a>";
    var page_usd_html = "<a target=\"_blank\" href=\""+usd_svg+"\"><img src=\""+usd_svg+"\"></a>";
    var page_color_html = "<a target=\"_blank\" href=\""+color_svg+"\"><img src=\""+color_svg+"\"></a>";

    $("#show_badge").html(page_html+" "+page_usd_html+" "+page_color_html);

    $("#html_source").val(html);
    $("#markup_source").val(markup);
    $("#restruct_source").val(restruct);
    $("#ascii_source").val(ascii);
    $("#render_img").removeClass('d-none');
    return false;
});