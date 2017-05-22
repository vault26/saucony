$(function(){
  $(document).scroll(function(){
    var scrollTop = $(this).scrollTop();
    var $nav = $('nav');
    if (scrollTop > 145) {
      $nav.addClass('fix-top');
    } else if (scrollTop < 40) {
      $nav.removeClass('fix-top');
    }
  });
});
