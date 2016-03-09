$(document).ready(function() {

  $.ajax({
    url: "/allstrokes"
  }).done(function(msg) {
    $('#foo').text(msg);
  });
})

