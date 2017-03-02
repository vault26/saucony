$(function(){
  $('form input').on('change', function(){
    $(this).parents('form').submit();
  });
});