

$(document).ready(function() {
  var filters = []
  filters['meal'] = []
  filters['hall'] = []
  filters['filter'] = []

  function toggleFilter(type, value) {
    console.log(type)
    console.log(filters[type])
    var index = filters[type].indexOf(value);
    if (index < 0)
      filters[type].push(value);
    else 
      filters[type].splice(index, 1);
  }

  function hasValue(type, value) {
    return filters[type].indexOf(value) != -1
  }

  function hasFilters() {
    return (filters['meal'].length + filters['hall'].length + 
      filters['filter'].length != 0)
  }

 $('.upvote').click(function() {
  var id = $(this).data('food-id')
  var vote = $('#' + id).find('.votes')
  vote.html(parseInt(vote.text())+1);
  $.post( "/api/inc/" + id);
});
 $('.downvote').click(function() {
  var id = $(this).data('food-id')
  var vote = $('#' + id).find('.votes')
  vote.html(parseInt(vote.text())-1);
  $.post("/api/dec/" + id);
});

 $('.circle').click(function() {

  var type = $(this).attr('class');
  var value = $(this).next('.filter_name').text();

  type = $.trim(type.replace('circle', '').replace('highlight', ''));

  value = value.split(' ').join('/'); // replace spaces with slashes for dining halls
  if (type == 'meal')
    value = value[0].toLowerCase();

  toggleFilter(type, value);
  console.log(filters[type]);

  // highlight border
  if ($(this).hasClass('highlight'))
    $(this).removeClass('highlight');
  else
    $(this).addClass('highlight');


  // reveal foods
  $('.food').each(function() {
    var valid = true;
    for (var ls in filters) {
      var valid_filt = (filters[ls].length == 0);
      for (var idx in filters[ls]) {
        valid_filt |= $(this).hasClass(filters[ls][idx]);
      }
      valid &= valid_filt;
    }

    if (valid)
      $(this).removeClass('hidden');
    else
      $(this).addClass('hidden');
  });
    
});
});
