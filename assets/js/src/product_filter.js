$(function(){
  $('#products-list input').on('change', function(){
    $(this).parents('form').submit();
  });
});
