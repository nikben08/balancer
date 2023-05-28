$(document).ready(function(){
    $("#createServer").click(function(){
        let server_address = $("#serverAdress").val();
        let server_label = $("#serverLabel").val();
        let server_weight = $("#serverWeight").val();
        let domain_id = $("#domainId").val();    
        
        if (server_address.length == 0){
          $("#serverAdress").css('border', '2px solid red');
        } 

        if (server_label.length == 0){
          $("#serverLabel").css('border', '2px solid red');
        } 

        if (server_weight.length == 0){
          $("#serverWeight").css('border', '2px solid red');
        } 

        if (domain_id.length == 0){
          $("#domainId").css('border', '2px solid red');
        } 

        if (server_weight.length != 0 && server_label.length != 0 && server_weight.length != 0){
          $.ajax({
              method: "POST",
              url: "/dashboard/create_server",
              data: JSON.stringify({
                "server_address": server_address,
                "server_label": server_label,
                "server_weight": server_weight,
                "domain_id": domain_id
              }),
              contentType: "application/json; charset=utf-8",
              traditional: true,
            })
              .done(function( msg ) {
                if (msg["code"] == 200) {
                  location.reload()
                }else{
                  $("#modal-error_box").append('<div class="alert alert-danger" role="alert">' + msg["message"] + '</div>')
                }
              });
        }

    });


    $(".delete_server-btn").click(function (){
      let server_id = $(this).data('server-id');
      if (confirm("bu sunucuyu silmek istediğinizden emin misiniz?") == true){
        $.ajax({
          method: "POST",
          url: "/dashboard/delete_server",
          data: JSON.stringify({
            "server_id": server_id,
          }),
          contentType: "application/json; charset=utf-8",
          traditional: true,
        })
          .done(function( msg ) {
            if (msg["code"] == 200) {
              location.reload();
            }else{
              alert(msg["message"]);
            }
          });
      }
    });



    $("#createDomain").click(function(){
      let domain_label = $("#domainLabel").val();

      if (domain_label.length == 0){
        $("#domainLabel").css('border', '2px solid red');
      } 

      if (domain_label.length != 0){
        $.ajax({
            method: "POST",
            url: "/dashboard/create_domain",
            data: JSON.stringify({
              "domain_label": domain_label,
            }),
            contentType: "application/json; charset=utf-8",
            traditional: true,
          })
            .done(function( msg ) {
              if (msg["code"] == 200) {
                location.reload()
              }else{
                $("#modal-error_box").append('<div class="alert alert-danger" role="alert">' + msg["message"] + '</div>')
              }
            });
      }

  });


  $(".delete_domain-btn").click(function (){
    let domain_id = $(this).data('domain-id');
    if (confirm("bu etki alanı silmek istediğinizden emin misiniz?") == true){
      $.ajax({
        method: "POST",
        url: "/dashboard/delete_domain",
        data: JSON.stringify({
          "domain_id": domain_id,
        }),
        contentType: "application/json; charset=utf-8",
        traditional: true,
      })
        .done(function( msg ) {
          if (msg["code"] == 200) {
            location.reload();
          }else{
            alert(msg["message"]);
          }
        });
    }
  });


  $("#domainSelect").on('change', function (){
      // Create a form element dynamically
    var form = $('<form>', {
      action: 'peers_manager',
      method: 'post'
    });

    // Create input fields and set their values
    $('<input>').attr({
      type: 'hidden',
      name: 'DomainId',
      value: this.value
    }).appendTo(form);

    // Append the form to the document body and submit it
    form.appendTo('body').submit();
  });

});