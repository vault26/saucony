$(function(){
  $('#mobile-menu button').click(function(){
    var $menus = $(this).next();
    if ($menus.is(':hidden')) {
      $(this).next().slideDown();
      $(this).addClass('active');
    } else {
      $(this).next().slideUp();
      $(this).removeClass('active');
    }
  });
});
