// animate.css
var ANIMATE = {
  animationEnd: 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend',
  addAnimation: function($component, animationClass) {
    $component.addClass(animationClass)
      .one(ANIMATE.animationEnd, function() {
        $component.removeClass(animationClass);
    });
  }
};


