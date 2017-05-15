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
    $errorMessage.hide();
    $successMessage.hide();

    // use vanilla element to get attached files
    fileInput = fileInput.get(0);
    if(fileInput.files.length == 0) {
      $(this).prop('disabled', false);
      return;
    }
    uploadTransferSlip(url, orderId, fileInput.files[0],
                       $successMessage, $errorMessage, $(this));
  });

  function uploadTransferSlip(url, orderId, file, $success, $error, $button) {
    var form = new FormData();
    form.append('orderId', orderId);
    form.append("Content-Type", file.type);
    form.append('file', file);
    axios
      .post(url, form)
      .then(function(res){
        $success.slideDown();
      })
      .catch(function(err){
        $error.slideDown();
        $button.prop('disabled', false);
      });
  }

});
