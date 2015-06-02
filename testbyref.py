__author__ = 'bjcheny'

# http://stackoverflow.com/questions/986006/how-do-i-pass-a-variable-by-reference

def try_to_change_list_content(the_list):
    print 'got', the_list
    the_list.append('four')
    print 'changed to', the_list

def try_to_change_list_reference(the_list):
    print 'got', the_list
    the_list = ['and', 'we', 'can', 'not', 'lie']
    print 'set to', the_list

def try_to_change_string_reference(the_string):
    print 'got', the_string
    the_string = 'In a kingdom by the sea'
    print 'set to', the_string

def main():
    outer_list = ['one', 'two', 'three']
    print 'before, outer_list = ', outer_list
    try_to_change_list_content(outer_list)
    print 'after, outer_list =', outer_list

    outer_list = ['we', 'like', 'proper', 'English']
    print 'before, outer_list =', outer_list
    try_to_change_list_reference(outer_list) # The the_list was a copy of the outer_list reference
    print 'after, outer_list =', outer_list


    outer_string = 'It was many and many a year ago'
    print 'before, outer_string =', outer_string
    try_to_change_string_reference(outer_string)
    print 'after, outer_string =', outer_string

if __name__ == '__main__':
    main()
