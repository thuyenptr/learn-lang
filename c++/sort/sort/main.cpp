//
//  main.cpp
//  sort
//
//  Created by Meteor on 8/27/19.
//  Copyright © 2019 Meteor. All rights reserved.
//

#include <iostream>
using namespace std;

void merge(int array[], int left, int mid, int right) {
    int nPart1 = mid - left + 1;
    int nPart2 = right - mid;
    
    int Part1[nPart1];
    int Part2[nPart2];
    
    // copy value in array to part 1
    for (int i = 0; i < nPart1; ++i) {
        Part1[i] = array[left + i];
    }
    
    // copy value in array to part 2
    for (int i = 0; i < nPart2; ++i) {
        Part2[i] = array[mid + i + 1];
    }
    
    // merge part 1 and part 2
    int i = 0, j = 0, k = left;
    while (i < nPart1 && j < nPart2) {
        if (Part1[i] <= Part2[j]) {
            array[k] = Part1[i];
            i++;
        } else {
            array[k] = Part2[j];
            j++;
        }
        k++;
    }
    
    while (i < nPart1) {
        array[k++] = Part1[i++];
    }
    
    while (j < nPart2) {
        array[k++] = Part2[j++];
    }
}

void merge_sort(int array[], int left, int right) {
    if (left < right) {
        int mid = (left + right) / 2;
        merge_sort(array, left, mid);
        merge_sort(array, mid + 1, right);
        merge(array, left, mid, right);
    }
}

void show_array(int array[], int n) {
    for (int i = 0; i < n; ++i) {
        cout << array[i] << " ";
    }
}

int main(int argc, const char * argv[]) {
    int array[] = {3,1,2,4,9,5,8,7,6};
    show_array(array, 9);
    cout << endl;
    merge_sort(array, 0, 8);
    show_array(array, 9);
    return 0;
}
