$(function(){
  var selectedModel = $('#stores #product-model').val();
  showModelOptions(selectedModel);

  $('#stores #product-model').change(function(){
    updateColor($(this).val());
  });

  function updateColor(model) {
    $productColorSelection = $('#stores #product-color');
    $productColorSelection.val('');
    $productColorSelection.find('option.color').hide();
    showModelOptions(model);
  }

  function showModelOptions(model) {
    $('#stores #product-color')
      .find('option.color[data-model="'+ model +'"]')
      .show();
  }

  $('#find-shoes-in-stores').click(function(){
    var $form = $('#stores form');
    var formHidden = $form.is(':hidden');
    if (formHidden) {
      $form.slideDown();
    } else {
      $form.slideUp();
    }
  });
});
