#include <iostream> // for standard I/O
#include <string>   // for strings
#include <iomanip>  // for controlling float print precision
#include <sstream>  // string to number conversion
#include <pthread.h>

#include "FishEyeCorrect.h"
#include  "opencv2/core/core.hpp"        // Basic OpenCV structures (cv::Mat, Scalar)
#include  "opencv2/imgproc/imgproc.hpp"  // Gaussian Blur
#include  "opencv2/highgui/highgui.hpp"// OpenCV window I/O

using namespace std;
using namespace cv;

int frameNumber = 0;
void  *thrd_func1(void *argv);
void  *thrd_func2(void *argv);
void  *thrd_func3(void *argv);
void  *thrd_func4(void *argv);
CFishEyeCorrect fishCorrect;
	
int main(int argc, char *argv[])
{
//	string path = "C:\\Users\\Administrator\\Desktop\\360TSR_2.mp4";
	VideoCapture cap(argv[1]);
	frameNumber = cap.get(CV_CAP_PROP_FRAME_COUNT);
	cout<<"frameNumber "<<frameNumber<<endl;
	if(!cap.isOpened())
	{
		cout<<"读取视频失败！"<<endl;
		return 0;
	}
//	CFishEyeCorrect fishCorrect;
	int flag = 0;
	namedWindow("Video0");
	namedWindow("Video1");
	namedWindow("Video2");
	namedWindow("Video3");
	namedWindow("Video4");
	for(int i = 0;i < frameNumber;++i)
	{
		Mat frame;
		pthread_t tid1,tid2,tid3,tid4;
     void *tret;
		clock_t t1 = clock();
	  //	double posFrameIndex = cap.get(CV_CAP_PROP_POS_FRAMES);
		//pDlg->m_ctrlProgress.SetPos(posFrameIndex);
		//CString text;
		//text.Format("%d",(int)posFrameIndex);
		//pDlg->SetDlgItemText(IDC_TEXT,_T(text));

		cap >> frame;
		if (!frame.data)
			{
				break;
		   }
		//clock_t start = clock();
		//resize(frame,frame,Size(frame.cols/2,frame.rows/2));
		//clock_t start = clock();
 //	pDlg->ShowImage(frame,IDC_VIDEO1);
		//flip(frame,frame,-1);
		//Mat temp;
		//resize(frame,temp,Size(241,151));
		//pDlg->ShowImage(CSize(frame.cols,frame.rows),24,frame.data,1);

		clock_t t2 = clock();
		cout <<"jiema :"<<(t2-t1)/1000.0<<endl;
		if (flag == 0)
		{
			fishCorrect.GetArea(frame);
			flag =1;
		}
		//Mat tempImg = frame(fishCorrect.correctArea);
		if (i>0)
			{
				// 等待线程tid1结束，线程函数返回值放在tret中
		     if (pthread_join(tid1,&tret)!=0){
		         printf("Join thread 1 error!\n");
		         exit(1);
		     }
		 
		    // printf("Thread 1 exit code: %d.\n",(int *)tret);
		     // 等待tid2结束，线程函数返回值放在tret中
		    if (pthread_join(tid2,&tret)!=0){
		         printf("Join thread 2 error!\n");
		        exit(1);
		     }
				 // 等待tid2结束，线程函数返回值放在tret中
		     if (pthread_join(tid3,&tret)!=0){
		         printf("Join thread 3 error!\n");
		         exit(1);
		     }	
		
		     if (pthread_join(tid4,&tret)!=0){
		         printf("Join thread 4 error!\n");
		         exit(1);
		     }	
			  clock_t end = clock();
			  // cout<<"fix is :"<<(end -start)/1000.0<<endl;
				imshow("Video0",fishCorrect.orgImage);
				imshow("Video1",fishCorrect.dstImage1);
				imshow("Video2",fishCorrect.dstImage2);
				imshow("Video3",fishCorrect.dstImage3);
				imshow("Video4",fishCorrect.dstImage4);
				waitKey(1);
		   	clock_t t3 = clock();
			  cout << "show is:"<<(t3 - end)/1000.0<<endl;
						
				}
		fishCorrect.orgImage  = frame(fishCorrect.correctArea);
		if (i ==0)
		{
			fishCorrect.CalculateParam(fishCorrect.orgImage.cols,fishCorrect.orgImage.rows);
			fishCorrect.FishImgIndex(fishCorrect.orgImage);
			//FIC.Flag = 0;
		}
		//clock_t start = clock();
		// SM.Standard_circle(tempImg,imgCircle);

	/*	Mat dstImg1;
		Mat dstImg2;
		Mat dstImg3;

		fishCorrect.ImageCorrect(tempImg,dstImg1,dstImg2,dstImg3);*/
		//FIC.ImageCorrect(frame,dstImg1,dstImg2,dstImg3);
		/*clock_t end = clock();
		CString time;
		time.Format("%d\n%d",end - middle,middle - start);
		pDlg->MessageBox(time);*/
		// writer << dstImg1;


		//double angle = -90;
		//double scale = 1;
		//Mat WarpImg1;
		//Mat WarpImg2;
		//Mat WarpImg3;
		////pDlg->RotationwarpAffine(dstImg1,WarpImg1,angle,scale);
		//pDlg->RotationwarpAffine(dstImg2,WarpImg2,angle,scale);
		//pDlg->RotationwarpAffine(dstImg3,WarpImg3,angle,scale);

		//pDlg->ShowImage(CSize(WarpImg1.cols,WarpImg1.rows),24,WarpImg1.data,2);
		//pDlg->ShowImage(CSize(dstImg2.cols,dstImg2.rows),24,dstImg2.data,3);
		//pDlg->ShowImage(CSize(WarpImg2.cols,WarpImg2.rows),24,WarpImg2.data,4);
		
		//tempImg.copyTo(fishCorrect.orgImage);
    //clock_t start = clock();

     // 创建线程tid1，线程函数thrd_func1
     if (pthread_create(&tid1,NULL,thrd_func1,NULL)!=0) {
         printf("Create thread 1 error!\n");
         exit(1);
     }
     // 创建线程tid2,线程函数thrd_func2
     if (pthread_create(&tid2,NULL,thrd_func2,NULL)!=0) {
         printf("Create thread 2 error!\n");
         exit(1);
     }
			// 创建线程tid2,线程函数thrd_func2
     if (pthread_create(&tid3,NULL,thrd_func3,NULL)!=0) {
         printf("Create thread 3 error!\n");
         exit(1);
     }

     if (pthread_create(&tid4,NULL,thrd_func4,NULL)!=0) {
         printf("Create thread 4 error!\n");
         exit(1);
       }
		
			clock_t t4 = clock();
			cout << "all is:"<<(t4 - t1)/1000.0<<endl;
	}
	cout<<"game over！"<<endl;
	return 0;
}

void  *thrd_func1(void *argv)
{
	//Mat destImage;
	fishCorrect.ImageCorrect(fishCorrect.orgImage,fishCorrect.dstImage1,fishCorrect.ImgIndex1);
	//namedWindow("Video1");
	//imshow("Video1",destImage);
	//cout<<"pthread:1"<<endl;
	pthread_exit(NULL); 
	}
	
void  *thrd_func2(void *argv)
{
	//Mat destImage;
	fishCorrect.ImageCorrect(fishCorrect.orgImage,fishCorrect.dstImage2,fishCorrect.ImgIndex2);
	//namedWindow("Video2");
	//imshow("Video2",destImage);
	//cout<<"pthread:2"<<endl;
	pthread_exit(NULL); 
	
	}

void  *thrd_func3(void *argv)
{
	//Mat destImage;
	fishCorrect.ImageCorrect(fishCorrect.orgImage,fishCorrect.dstImage3,fishCorrect.ImgIndex3);
	//namedWindow("Video3");
	//imshow("Video3",destImage);
       // cout<<"pthread:3"<<endl;
	pthread_exit(NULL); 
	}

void  *thrd_func4(void *argv)
{
	//Mat destImage;
	fishCorrect.ImageCorrect(fishCorrect.orgImage,fishCorrect.dstImage4,fishCorrect.ImgIndex4);
	//namedWindow("Video3");
	//imshow("Video3",destImage);
       // cout<<"pthread:3"<<endl;
	pthread_exit(NULL); 
	}
