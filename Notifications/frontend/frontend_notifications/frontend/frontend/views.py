from django.shortcuts import render, redirect
from django.http import HttpResponse, HttpResponseBadRequest
from django.core.paginator import Paginator, EmptyPage, PageNotAnInteger

# Create your views here.

from django.shortcuts import render
from django.http import Http404

from django.contrib.auth.decorators import login_required

from django.core.files.storage import FileSystemStorage


def index(request):

    return render(request, 'html/test.html')