$(function(){
  $('#payment-upload button').click(function(e){
    e.preventDefault();
    $(this).prop('disabled', true);
    var $form = $(this).parent();
    var url = $form.prop('action');
    var orderId = $form.find('input[name="orderId"]').val();
    var fileInput = $form.find('input[type="file"]');
    var $errorMessage = $form.find('.error');
    var $successMessage = $form.find('.success');
    var $loading = $form.find('.loading');
    $errorMessage.hide();
    $successMessage.hide();
    $loading.show();

    // use vanilla element to get attached files
    fileInput = fileInput.get(0);
    if(fileInput.files.length == 0) {
      $(this).prop('disabled', false);
      $loading.hide();
      return;
    }
    uploadTransferSlip(url, orderId, fileInput.files[0],
                       $successMessage, $errorMessage, $(this), $loading);
  });

  function uploadTransferSlip(
    url, orderId, file, $success, $error, $button, $loading) {

    var form = new FormData();
    form.append('orderId', orderId);
    form.append("Content-Type", file.type);
    form.append('file', file);
    axios
      .post(url, form)
      .then(function(res){
        $success.slideDown();
        $loading.hide();
      })
      .catch(function(err){
        $error.slideDown();
        $button.prop('disabled', false);
        $loading.hide();
      });
  }

});
