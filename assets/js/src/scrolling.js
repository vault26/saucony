$(function(){
  $(window).scroll(function(){
    if (window.scrollY > 140) {
      $('#cart').addClass('stick-top');
    } else {
      $('#cart').removeClass('stick-top');
    }
  });
});
