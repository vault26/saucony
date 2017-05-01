$(function(){
  var selectedModel = $('#stores #product-model').val();
  updateColor(selectedModel);

  $('#stores #product-model').change(function(){
    updateColor($(this).val());
  });

  function updateColor(model) {
    $productColorSelection = $('#stores #product-color');
    $productColorSelection.val('');

    $productColorSelection.find('option.color').hide();
    $modelColorOptions = $productColorSelection
      .find('option.color[data-model="'+ model +'"]');
    $modelColorOptions.show();
  }
});
