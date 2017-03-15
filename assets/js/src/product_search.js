$(function(){
  $('#clear-search').click(function(){
    var href = window.location.href;
    href = href.replace(/(query=.+?)(&|$)/, '$2');
    window.location.href = href;
  });
});
