#include <stdio.h>
#include <stdlib.h>

// -------------------------------------------------------------

typedef struct _MemberInfo {
	const char* tag;
	void* addr;
} MemberInfo;

typedef struct _TypeInfo {
	MemberInfo* members;
} TypeInfo;

void* MemberFind(TypeInfo* ti, const char* tag) {
	size_t n = 0;
	while (ti->members[n].tag != NULL) {
		if (strcmp(ti->members[n].tag, tag) == 0) {
			return ti->members[n].addr;
		}
		n++;
	}
	return NULL;
}

// -------------------------------------------------------------

typedef struct _InterfaceInfo {
	const char** tags;
} InterfaceInfo;

size_t MemberCount(InterfaceInfo* intf) {
	size_t n = 0;
	while (intf->tags[n] != NULL) {
		n++;
	}
	return n;
}

// -------------------------------------------------------------

typedef struct _IReadWriterTbl {
	InterfaceInfo* inter;
	TypeInfo* type;
	int (*Read)(void* this, char* buf, int cb);
	int (*Write)(void* this, char* buf, int cb);
} IReadWriterTbl;

typedef struct _IReadWriter {
	IReadWriterTbl* tab;
	void* data;
} IReadWriter;

const char* g_Tags_IReadWriter[] = {
	"Read(*char,int)int",
	"Write(*char,int)int",
	NULL
};

InterfaceInfo g_InterfaceInfo_IReadWriter = {
	g_Tags_IReadWriter
};

// -------------------------------------------------------------

typedef struct _IWriterTbl {
	InterfaceInfo* inter;
	TypeInfo* type;
	int (*Write)(void* this, char* buf, int cb);
} IWriterTbl;

typedef struct _IWriter {
	IWriterTbl* tab;
	void* data;
} IWriter;

const char* g_Tags_IWriter[] = {
	"Write(*char,int)int",
	NULL
};

InterfaceInfo g_InterfaceInfo_IWriter = {
	g_Tags_IWriter
};

// -------------------------------------------------------------

typedef struct _A {
	int a;
} A;

int A_Read(A* this, char* buf, int cb) {
	printf("A_Read: %d\n", this->a);
	return cb;
}

int A_Write(A* this, char* buf, int cb) {
	printf("A_Write: %d\n", this->a);
	return cb;
}

MemberInfo g_Members_A[] = {
	{ "Read(*char,int)int", A_Read },
	{ "Write(*char,int)int", A_Write },
	{ NULL, NULL }
};

TypeInfo g_TypeInfo_A = {
	g_Members_A
};

A* NewA(int params) {
	printf("NewA: %d\n", params);
	A* this = (A*)malloc(sizeof(A));
	this->a = params;
	return this;
}

// -------------------------------------------------------------

typedef struct _B {
	A base;
} B;

int B_Write(B* this, char* buf, int cb) {
	printf("B_Write: %d\n", this->base.a);
	return cb;
}

void B_Foo(B* this) {
	printf("B_Foo: %d\n", this->base.a);
}

MemberInfo g_Members_B[] = {
	{ "Read(*char,int)int", A_Read },
	{ "Write(*char,int)int", B_Write },
	{ "Foo()", B_Foo },
	{ NULL, NULL }
};

TypeInfo g_TypeInfo_B = {
	g_Members_B
};

B* NewB(int params) {
	printf("NewB: %d\n", params);
	B* this = (B*)malloc(sizeof(B));
	this->base.a = params;
	return this;
}

// -------------------------------------------------------------

IWriterTbl g_Itbl_IWriter_B = {
	&g_InterfaceInfo_IWriter,
	&g_TypeInfo_B,
	(int (*)(void* this, char* buf, int cb))B_Write
};

// -------------------------------------------------------------

typedef struct _ITbl {
	InterfaceInfo* inter;
	TypeInfo* type;
	//...
} ITbl;

ITbl* MakeItbl(InterfaceInfo* intf, TypeInfo* ti) {
	size_t i, n = MemberCount(intf);
	ITbl* dest = (ITbl*)malloc(n * sizeof(void*) + sizeof(ITbl));
	void** addrs = (void**)(dest + 1);
	for (i = 0; i < n; i++) {
		addrs[i] = MemberFind(ti, intf->tags[i]);
		if (addrs[i] == NULL) {
			free(dest);
			return NULL;
		}
	}
	dest->inter = intf;
	dest->type = ti;
	return dest;
}

// -------------------------------------------------------------

int main() {
	B* unnamed = NewB(10);
	IWriter p = {
		&g_Itbl_IWriter_B,
		unnamed
	};
	IReadWriter p2 = {
		(IReadWriterTbl*)MakeItbl(&g_InterfaceInfo_IReadWriter, p.tab->type),
		p.data
	};
	int ok = (p2.tab != NULL);
	p.tab->Write(p.data, NULL, 10);
	if (ok) {
		p2.tab->Read(p2.data, NULL, 10);
	}
	return 0;
}

// -------------------------------------------------------------

